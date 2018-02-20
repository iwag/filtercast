package io.github.iwag.filtercast.repositories;


import io.github.iwag.filtercast.models.RSSEntity;

import java.util.List;
import java.util.Optional;

public interface TaskService {
    Optional<Long> createTask(RSSEntity task);

    RSSEntity readEntity(Long taskId);

    void updateEntity(RSSEntity task);

    void deleteTask(Long taskId);

    List<RSSEntity> listTasks(String startCursor);
}
