/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

export default function (visitable, submitable, deletable, cancelable, clickable) {
  return submitable(
    cancelable(
      deletable({
        visit: visitable(['/:dc/acls/:acl', '/:dc/acls/create']),
        use: clickable('[data-test-use]'),
        confirmUse: clickable('[data-test-confirm-use]'),
      })
    ),
    'main'
  );
}
