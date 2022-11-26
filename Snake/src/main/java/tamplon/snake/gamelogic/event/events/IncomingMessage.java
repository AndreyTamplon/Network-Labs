package tamplon.snake.gamelogic.event.events;

import tamplon.snake.network.message.AddressedMessage;
import tamplon.snake.gamelogic.event.Event;

public class IncomingMessage implements Event
{
    public AddressedMessage message;

    public IncomingMessage(AddressedMessage message)
    {
        this.message = message;
    }
}
