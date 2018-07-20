/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

/*
author: github.com/jasonkylelol
date: 2018-07-20
*/

package hkex_news_fetcher

import (
	"github.com/jasonkylelol/hkex_news_fetcher/fetcher"
)

type AnnInfo struct {
	Guid      string
	FileTitle string
	FileLink  string
	StockID   string
	StockName string
	TimeStr   string
	Timestamp int64
	Content   string
}

func LatestAnn() ([]AnnInfo, error) {
	results := []AnnInfo{}
	infos := []fetcher.AnnInfo{}
	if err := fetcher.FetchLatestAnn(&infos); err != nil {
		return results, err
	}
	for _, v := range infos {
		r := AnnInfo{
			Guid:      v.Guid,
			FileTitle: v.FileTitle,
			FileLink:  v.FileLink,
			StockID:   v.StockID,
			StockName: v.StockName,
			TimeStr:   v.TimeStr,
			Timestamp: v.Timestamp,
			Content:   v.Content,
		}
		results = append(results, r)
	}
	return results, nil
}

func SearchAnn(stock string, start_timestamp, end_timestamp int64) ([]AnnInfo, error) {
	results := []AnnInfo{}
	infos := []fetcher.AnnInfo{}
	for _, v := range infos {
		r := AnnInfo{
			Guid:      v.Guid,
			FileTitle: v.FileTitle,
			FileLink:  v.FileLink,
			StockID:   v.StockID,
			StockName: v.StockName,
			TimeStr:   v.TimeStr,
			Timestamp: v.Timestamp,
			Content:   v.Content,
		}
		results = append(results, r)
	}
	return results, nil
}