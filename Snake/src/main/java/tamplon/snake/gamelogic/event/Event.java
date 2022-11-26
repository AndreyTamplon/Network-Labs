package tamplon.snake.gamelogic.event;

public interface Event
{
    @SuppressWarnings("unchecked")
    default <T extends Event> T get()
    {
        return (T) this;
    }
}
