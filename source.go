/*  Copyright 2012, mokasin
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package main

// TODO Documenation

type TrackTags struct {
	Path    string
	Title   string
	Artist  string
	Album   string
	Comment string
	Genre   string
	Year    uint
	Track   uint
	Bitrate uint
	Length  uint
}

type TrackInfo interface {
	Path() string
	Mtime() int64
	Tags() (*TrackTags, error)
}

type TrackSource interface {
	Crawl(tracks chan<- TrackInfo, done chan<- bool)
}