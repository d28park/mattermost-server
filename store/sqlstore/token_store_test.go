// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/store/storetest"
)

func TestTokenStore(t *testing.T) {
	StoreTest(t, storetest.TestTokenStore)
}
