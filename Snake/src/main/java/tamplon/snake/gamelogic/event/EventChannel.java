package tamplon.snake.gamelogic.event;

@FunctionalInterface
public interface EventChannel
{
    void submit(Event event) throws InterruptedException;
}
