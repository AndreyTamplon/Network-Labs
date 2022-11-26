package tamplon.snake.gamelogic.event.events;

import tamplon.snake.network.message.AddressedMessage;
import tamplon.snake.gamelogic.event.Event;

public class Announcement implements Event
{
    public final AddressedMessage message;

    public Announcement(AddressedMessage message)
    {
        this.message = message;
    }
}
