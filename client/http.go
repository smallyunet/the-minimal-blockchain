package client

//func GetLatestBlock() {
//	ticker := time.NewTicker(time.Duration(blockTime) * time.Millisecond)
//	for {
//		select {
//		case <-ticker.C:
//			height, err := storage.GetHeight()
//			if err != nil {
//				log.Println(err)
//			}
//			for k, _ := range RouteTable {
//				requestURL := fmt.Sprintf("http://%s/get/%d", k, height+1)
//				get, err := http.Get(requestURL)
//				if err != nil {
//					continue
//				}
//				var b *block.Block
//				all, err := io.ReadAll(get.Body)
//				if err != nil {
//					continue
//				}
//				err = json.Unmarshal(all, &b)
//				if err != nil {
//					continue
//				}
//				pool.PushBlockToPool(b)
//			}
//		}
//	}
//}
