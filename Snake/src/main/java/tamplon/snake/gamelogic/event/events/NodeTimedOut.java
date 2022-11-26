package tamplon.snake.gamelogic.event.events;

import tamplon.snake.gamelogic.event.Event;

import java.net.InetSocketAddress;

public class NodeTimedOut implements Event
{

    public InetSocketAddress nodeAddress;

    public NodeTimedOut(InetSocketAddress nodeAddress)
    {
        this.nodeAddress = nodeAddress;
    }
}
