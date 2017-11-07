/*
 * Copyright 2017 Kopano and its licensors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, version 3,
 * as published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package kcc

// A LogonResponse holds tthe returned data of a SOAP logon request.
type LogonResponse struct {
	Er         uint64 `xml:"er"`
	SessionID  uint64 `xml:"ulSessionId"`
	ServerGUID string `xml:"sServerGuid"`
}

// A LogoffResponse holds the returned data of a SOAP logoff request.
type LogoffResponse struct {
	Er uint64 `xml:"er"`
}

// A ResolveUserResponse holds the returned data of a SOAP request which
// retruns a user's ID details.
type ResolveUserResponse struct {
	Er          uint64 `xml:"er"`
	ID          uint64 `xml:"ulUserId"`
	UserEntryID string `xml:"sUserId"`
}

// A GetUserResponse holds the returned data of a SOAP request which fetches
// user detail meta data.
type GetUserResponse struct {
	Er   uint64 `xml:"er"`
	User *User  `xml:"lpsUser"`
}

// A User represents the meta data of a user as stored by Kopano server.
type User struct {
	ID          uint64 `xml:"ulUserId" json:"ulUserID"`
	Username    string `xml:"lpszUsername" json:"lpszUsername"`
	MailAddress string `xml:"lpszMailAddress" json:"lpszMailAddress"`
	FullName    string `xml:"lpszFullName" json:"lpszFullName"`
	IsAdmin     uint64 `xml:"ulIsAdmin" json:"ulIsAdmin"`
	IsNonActive uint64 `xml:"ulIsNonActive" json:"ulIsNonActive"`
	UserEntryID string `xml:"sUserId" json:"sUserId"`
}
