package tamplon.snake.gamelogic.event;

@FunctionalInterface
public interface EventHandler
{
    void handle(Event message) throws Exception;
}
