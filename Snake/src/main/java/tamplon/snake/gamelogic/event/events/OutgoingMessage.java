package tamplon.snake.gamelogic.event.events;

import tamplon.snake.network.message.AddressedMessage;
import tamplon.snake.gamelogic.event.Event;

public final class OutgoingMessage implements Event
{
    public AddressedMessage message;

    public OutgoingMessage(AddressedMessage message)
    {
        this.message = message;
    }
}
