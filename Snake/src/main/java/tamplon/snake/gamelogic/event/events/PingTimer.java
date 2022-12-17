package tamplon.snake.gamelogic.event.events;

import tamplon.snake.gamelogic.event.Event;

import java.net.InetSocketAddress;

public class PingTimer implements Event
{
    public InetSocketAddress who;

    public PingTimer(InetSocketAddress who)
    {
        this.who = who;
    }
}
