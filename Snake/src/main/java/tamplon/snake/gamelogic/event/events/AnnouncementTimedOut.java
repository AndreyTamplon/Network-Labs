package tamplon.snake.gamelogic.event.events;

import tamplon.snake.gamelogic.event.Event;

import java.net.InetSocketAddress;

public class AnnouncementTimedOut implements Event
{

    public InetSocketAddress fromAddress;

    public AnnouncementTimedOut(InetSocketAddress fromAddress)
    {
        this.fromAddress = fromAddress;
    }
}
