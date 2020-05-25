# WebSocket内部設計書
最終更新日 : 2020 05 18  

## チャットルーム内画面

`/chatroom`にwebsocketで接続。  

### ルーム::ルーム参加リクエスト

- ルーム参加リクエストを受け取った際の処理の手順を以下に記載します。
	1. Socket.ioのルームにルームIDでジョイン処理を行います。
	2. 'notify_entryイベント' で、クライアントに通知します。

- シーケンス

	![参加](https://github.com/taniwhy/mochi-match-rest/blob/doc/out/websocket/%E3%83%AB%E3%83%BC%E3%83%A0%E5%8F%82%E5%8A%A0/%E3%83%AB%E3%83%BC%E3%83%A0%E5%8F%82%E5%8A%A0.png?raw=true)

__On  Event Name__  : join_req  
__On Data Format__ :  

	{
		'user': {
			'id': string,
			'user_name': string,
			'icon': string
			'favarate_game': [
				{
					'game_title_id': string,
					'game_title': string
				}
			],
		},
		'room_id': string,
		'timestamp': int
	}

### ルーム::ルーム参加通知

__Emit Event Name__ : notify_entry  
__Emit Data Format__ :  

	{
		'user': {
			'id': string,
			'user_name': string,
			'icon': string
			'favarate_game': [
				{
					'game_title_id': string,
					'game_title': string
				}
			],
		},
		'timestamp': int
	}

### ルーム::切断時処理

- Websocket接続が切断された際の処理の手順を以下に記載します。
	1. クライアントからdissconnectイベントを受け取ります。
	2. 保持しているroom_idとuser_idでAPIサーバのルーム退室APIを叩きます。
	3. 'notify_leaveイベント'で、クライアントに通知します。

- シーケンス

	![切断](https://github.com/taniwhy/mochi-match-rest/blob/doc/out/websocket/%E3%83%AB%E3%83%BC%E3%83%A0%E9%80%80%E5%AE%A4/%E3%83%AB%E3%83%BC%E3%83%A0%E9%80%80%E5%AE%A4.png?raw=true)

__On Event Name__ : disconnect
__On Data Format__ :  

	{}

### ルーム::ルーム退室通知

__Emit Event Name__ : notify_leave
__Emit Data Format__ :  

	{
		'user_id': string,
		'timestamp': int
	}

### エラー::エラー共通

__Emit Event Name__ : error  
__Emit Data Format__ :  

	{
		'message': string,
		'code': string,
		'type': string
	}
