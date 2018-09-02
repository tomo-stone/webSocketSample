// オブジェクト生成
var socket = new WebSocket("ws://"+ location.hostname +"/ws");
// 接続
socket.addEventListener("open", function(e){
  document.getElementById("status").innerHTML = "Connection is opened";
  socket.send("Connected to crient");
});

// 切断
window.unload = function(){
  socket.send(".");
  socket.close();
};

// メッセージ
socket.addEventListener("message", function(e){
  document.getElementById("clock").innerHTML = "<p>"+e.data+"</p>" ;
  socket.send("alive");
  count++;
});

// 更新カウンタ
var count=0;

var counter = function(){
  document.getElementById("count").innerHTML = "<p>count: "+count+"/sec.<p>" ;
  count=0;
}

setInterval(counter,1000);
