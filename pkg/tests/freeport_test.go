/*
 * Copyright (C) 2019 IsardVDI
 * Authors: IsardVDI Authors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package tests_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/brainupdaters/drlm-common/pkg/tests"

	"github.com/stretchr/testify/assert"
)

func TestGetFreePort(t *testing.T) {
	assert := assert.New(t)

	port := tests.GetFreePort(t)

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	assert.Nil(err)

	defer l.Close()
}
