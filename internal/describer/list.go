package describer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/heptio/developer-dash/pkg/store"
	"github.com/heptio/developer-dash/pkg/view/component"
)

// List describes a list of objects.
type List struct {
	*base

	path           string
	title          string
	listType       func() interface{}
	objectType     func() interface{}
	objectStoreKey store.Key
	isClusterWide  bool
}

// NewList creates an instance of List.
func NewList(p, title string, objectStoreKey store.Key, listType, objectType func() interface{}, isClusterWide bool) *List {
	return &List{
		path:           p,
		title:          title,
		base:           newBaseDescriber(),
		objectStoreKey: objectStoreKey,
		listType:       listType,
		objectType:     objectType,
		isClusterWide:  isClusterWide,
	}
}

// Describe creates content.
func (d *List) Describe(ctx context.Context, prefix, namespace string, options Options) (component.ContentResponse, error) {
	if options.Printer == nil {
		return EmptyContentResponse, errors.New("object list Describer requires a printer")
	}

	// Pass through selector if provided to filter objects
	var key = d.objectStoreKey // copy
	key.Selector = options.LabelSet

	if d.isClusterWide {
		namespace = ""
	}

	objects, err := options.LoadObjects(ctx, namespace, options.Fields, []store.Key{key})
	if err != nil {
		return EmptyContentResponse, err
	}

	list := component.NewList(d.title, nil)

	listType := d.listType()

	v := reflect.ValueOf(listType)
	f := reflect.Indirect(v).FieldByName("Items")

	// Convert unstructured objects to typed runtime objects
	for _, object := range objects {
		item := d.objectType()
		if err := scheme.Scheme.Convert(object, item, nil); err != nil {
			return EmptyContentResponse, err
		}

		if err := copyObjectMeta(item, object); err != nil {
			return EmptyContentResponse, err
		}

		newSlice := reflect.Append(f, reflect.ValueOf(item).Elem())
		f.Set(newSlice)
	}

	listObject, ok := listType.(runtime.Object)
	if !ok {
		return EmptyContentResponse, errors.Errorf("expected list to be a runtime object. It was a %T",
			listType)
	}

	viewComponent, err := options.Printer.Print(ctx, listObject, options.PluginManager())
	if err != nil {
		return EmptyContentResponse, err
	}

	if viewComponent != nil {
		if table, ok := viewComponent.(*component.Table); ok {
			list.Add(table)
		} else {
			list.Add(viewComponent)
		}
	}

	return component.ContentResponse{
		Components: []component.Component{list},
	}, nil
}

// PathFilters returns path filters for this Describer.
func (d *List) PathFilters() []PathFilter {
	return []PathFilter{
		*NewPathFilter(d.path, d),
	}
}