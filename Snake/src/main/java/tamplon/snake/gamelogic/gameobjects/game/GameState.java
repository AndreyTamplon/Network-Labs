package tamplon.snake.gamelogic.gameobjects.game;

import tamplon.snake.gamelogic.gameobjects.Utility;
import tamplon.snake.gamelogic.gameobjects.config.Config;
import tamplon.snake.gamelogic.gameobjects.player.PlayerInfo;
import tamplon.snake.gamelogic.gameobjects.snake.SnakeState;
import tamplon.snake.gamelogic.gamefield.Coordinates;
import tamplon.snake.gamelogic.gamefield.UnboundedFixedPoint;
import lombok.val;
import me.ippolitov.fit.snakes.SnakesProto;
import org.jetbrains.annotations.NotNull;

import java.util.ArrayList;

public interface GameState
{

    int getStateOrder();

    @NotNull Iterable<@NotNull SnakeState> getSnakes();

    @NotNull Iterable<@NotNull Coordinates> getFood();

    @NotNull Iterable<@NotNull PlayerInfo> getPlayers();

    @NotNull Config getConfig();

    static @NotNull GameState fromMessage(final @NotNull SnakesProto.GameStateOrBuilder state)
    {
        val config = state.getConfig();
        val playersList = state.getPlayers();
        val stateOrder = state.getStateOrder();

        val snakes = new ArrayList<@NotNull SnakeState>();
        for (int i = 0; i < state.getSnakesCount(); i += 1)
        {
            snakes.add(SnakeState.fromMessage(state.getSnakes(i)));
        }
        val food = new ArrayList<@NotNull Coordinates>();
        for (int i = 0; i < state.getFoodsCount(); i += 1)
        {
            val point = state.getFoods(i);
            food.add(new UnboundedFixedPoint(point.getX(), point.getY()));
        }
        val players = new ArrayList<@NotNull PlayerInfo>();
        for (int i = 0; i < playersList.getPlayersCount(); i += 1)
        {
            val player = playersList.getPlayers(i);
            players.add(PlayerInfo.fromMessage(player));
        }

        return new Game(
                stateOrder,
                snakes,
                food,
                players,
                Config.fromMessage(config));
    }

    default @NotNull SnakesProto.GameState toMessage()
    {
        val builder = SnakesProto.GameState.newBuilder()
                .setStateOrder(this.getStateOrder())
                .setConfig(this.getConfig().toMessage());
        val players = SnakesProto.GamePlayers.newBuilder();
        this.getPlayers().forEach(it -> players.addPlayers(it.toMessage()));
        builder.setPlayers(players);
        this.getSnakes().forEach(it -> builder.addSnakes(it.toMessage()));
        this.getFood().forEach(it -> builder.addFoods(Utility.coordinates(it)));
        return builder.build();
    }
}
