/*
 * Copyright (c) 2020 the Octant contributors. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class ElectronService {
  constructor() {}

  /**
   * returns true if electron is detected
   */
  isElectron(): boolean {
    return (
      process && process.versions && process.versions.electron !== undefined
    );
  }
}
