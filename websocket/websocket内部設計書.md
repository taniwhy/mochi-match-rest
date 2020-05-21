# WebSocket内部設計書
最終更新日 : 2020 05 18  

## チャットルーム内画面

`/chatroom`にwebsocketで接続。  

### ルーム::ルーム参加リクエスト

- ルーム参加リクエストを受け取った際の処理の手順を以下に記載します。
	1. リクエストに含まれるルームIDとユーザーIDを用いて、APIサーバのルーム参加APIを叩きます。
	2. 成功レスポンスを受け取った後、socket.ioのルームにルームIDでジョイン処理を行います。
	3. 'notify_joinイベント', 'join_resイベント'で、クライアントに通知します。

- シーケンス

	![参加](https://github.com/taniwhy/mochi-match-rest/blob/doc/out/websocket/%E3%83%AB%E3%83%BC%E3%83%A0%E5%8F%82%E5%8A%A0/%E3%83%AB%E3%83%BC%E3%83%A0%E5%8F%82%E5%8A%A0.png?raw=true)

- エラー
  - APIサーバからのレスポンスがエラーだった際は、'errorイベント' でクライアントに通知します。  

__On  Event Name__  : join_req  
__On Data Format__ :  

	{
		'details': {
			'user_id': int,
			'room_id': int
		}
	}

### ルーム::ルーム参加レスポンス

__Emit Event Name__ : join_res
__Emit Data Format__ :  

	{
		
	}


### ルーム::ルーム参加通知

__Emit Event Name__ : notify_join
__Emit Data Format__ :  

	{
		
	}

### ルーム::切断時処理

- Websocket接続が切断された際の処理の手順を以下に記載します。
	1. クライアントからdissconnectイベントを受け取ります。
	2. 保持しているroom_idとuser_idでAPIサーバのルーム退室APIを叩きます。
	3. 'notify_leaveイベント'で、クライアントに通知します。

- シーケンス

	![切断]()

__On Event Name__ : disconnect
__On Data Format__ :  
	{
		
	}

### ルーム::ルーム退室通知

__Emit Event Name__ : notify_leave
__Emit Data Format__ :  

	{
		
	}

### エラー::エラー共通

__Emit Event Name__ : error  
__Emit Data Format__ :  

	{
		'message': string,
		'code': string,
		'type': string
	}
