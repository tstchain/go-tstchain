// Copyright 2015 The go-tstereum Authors
// This file is part of the go-tstereum library.
//
// The go-tstereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-tstereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-tstereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the downloader.

package downloader

import (
	"github.com/tstchain/go-tstchain/metrics"
)

var (
	headerInMeter      = metrics.NewRegisteredMeter("tst/downloader/headers/in", nil)
	headerReqTimer     = metrics.NewRegisteredTimer("tst/downloader/headers/req", nil)
	headerDropMeter    = metrics.NewRegisteredMeter("tst/downloader/headers/drop", nil)
	headerTimeoutMeter = metrics.NewRegisteredMeter("tst/downloader/headers/timeout", nil)

	bodyInMeter      = metrics.NewRegisteredMeter("tst/downloader/bodies/in", nil)
	bodyReqTimer     = metrics.NewRegisteredTimer("tst/downloader/bodies/req", nil)
	bodyDropMeter    = metrics.NewRegisteredMeter("tst/downloader/bodies/drop", nil)
	bodyTimeoutMeter = metrics.NewRegisteredMeter("tst/downloader/bodies/timeout", nil)

	receiptInMeter      = metrics.NewRegisteredMeter("tst/downloader/receipts/in", nil)
	receiptReqTimer     = metrics.NewRegisteredTimer("tst/downloader/receipts/req", nil)
	receiptDropMeter    = metrics.NewRegisteredMeter("tst/downloader/receipts/drop", nil)
	receiptTimeoutMeter = metrics.NewRegisteredMeter("tst/downloader/receipts/timeout", nil)

	stateInMeter   = metrics.NewRegisteredMeter("tst/downloader/states/in", nil)
	stateDropMeter = metrics.NewRegisteredMeter("tst/downloader/states/drop", nil)
)
