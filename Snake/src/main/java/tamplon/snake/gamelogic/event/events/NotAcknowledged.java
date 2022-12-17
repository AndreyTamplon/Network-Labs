package tamplon.snake.gamelogic.event.events;

import tamplon.snake.network.message.AddressedMessage;
import tamplon.snake.gamelogic.event.Event;

public class NotAcknowledged implements Event
{
    public AddressedMessage message;

    public NotAcknowledged(AddressedMessage message)
    {
        this.message = message;
    }
}
