# WebSocket内部設計書
最終更新日 : 2020 05 17  

## チャットルーム内画面

`/chatroom`にwebsocketで接続。  

### ルーム::ルーム参加リクエスト

- ルーム参加リクエストを受け取った際の処理の手順を以下に記載します。
	1. リクエストに含まれるルームIDとユーザーIDを用いて、APIサーバのルーム参加APIを叩きます。
	2. 成功レスポンスを受け取った後、socket.ioのルームにルームIDでジョインさせます。
	3. 'notify_joinイベント', 'join_resイベント'で、クライアントに通知します。

- シーケンス

	![参加](https://github.com/taniwhy/mochi-match-rest/blob/doc/out/websocket/%E3%83%AB%E3%83%BC%E3%83%A0%E5%8F%82%E5%8A%A0/%E3%82%B7%E3%83%BC%E3%82%B1%E3%83%B3%E3%82%B9%E5%9B%B3.png?raw=true)

- エラー
  - APIサーバからのレスポンスがエラーだった際は、'errorイベント' でクライアントに通知します。 

__Emit  Event Name__  : join_req  
__Emit Data Format__ :  

	{
		'common' : {
			'type' : string
		},
		'details' : {
			user_id : string,
			room_id : string
		}
	}

### ルーム::ルーム参加通知

__On Event Name__ : notify_join
__On Data Format__ :  

	{
		
	}

### エラー::エラー共通

__Emit Event Name__ : error  
__Emit Data Format__ :  

	{
		'message' : string,
		'code' : string,
		'type' : string
	}
