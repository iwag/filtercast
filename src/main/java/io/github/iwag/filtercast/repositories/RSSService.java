package io.github.iwag.filtercast.repositories;


import io.github.iwag.filtercast.models.RSSEntity;

import java.util.List;
import java.util.Optional;

public interface RSSService {
    Optional<Long> createEntity(RSSEntity entity);

    RSSEntity readEntity(Long id);

    void updateEntity(RSSEntity entity);

    void deleteEntity(Long id);

    List<RSSEntity> listEntities(String startCursor);
}
