# fsm
Finite State Machine for golang </br>
所有状态机事件都异步处理，实现方式是放到事件队列处理。</br>
新建一个状态机，同时也会新建事件队列以及一个goruntine。 </br>
