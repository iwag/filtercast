package io.github.iwag.filtercast.repositories;

import com.google.appengine.api.datastore.*;
import io.github.iwag.filtercast.models.RSSEntity;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Optional;

@Service
public class DatastoreService implements RSSService {

    private final static String KIND = "RSS";

    @Autowired
    com.google.appengine.api.datastore.DatastoreService datastore;

    public RSSEntity entityToTask(Entity entity) {
        return new RSSEntity.Builder()
                .history((String) entity.getProperty(RSSEntity.HISTORY))
                .id(entity.getKey().getId())
                .url((String) entity.getProperty(RSSEntity.URL))
                .build();
    }

    @Override
    public Optional<Long> createEntity(RSSEntity entity) {
        Entity incTaskEntity = new Entity(KIND);
        incTaskEntity.setProperty(RSSEntity.HISTORY, entity.getHistory());
        incTaskEntity.setProperty(RSSEntity.URL, entity.getUrl());
        incTaskEntity.setProperty(RSSEntity.PUBLISHWAY, entity.getHistory());
        incTaskEntity.setProperty(RSSEntity.CREATED, entity.getCreated());
        incTaskEntity.setProperty(RSSEntity.UPDATED, entity.getUpdated());
        incTaskEntity.setProperty(RSSEntity.LATESTDATE, entity.getLatestdate());
        incTaskEntity.setProperty(RSSEntity.DURATION, entity.getDuration());

        Key k = datastore.put(incTaskEntity);
        return Optional.of(k.getId());
    }

    @Override
    public RSSEntity readEntity(Long id) {
        Entity entity = null;
        try {
            entity = datastore.get(KeyFactory.createKey(KIND, id));
        } catch (EntityNotFoundException e) {
            return null;
        }
        return entityToTask(entity);
    }

    @Override
    public void updateEntity(RSSEntity entity1) {
        Key key = KeyFactory.createKey(KIND, entity1.getId());
        Entity entity = new Entity(key);
        entity.setProperty(RSSEntity.HISTORY, entity1.getHistory());
        entity.setProperty(RSSEntity.URL, entity1.getUrl());
        entity.setProperty(RSSEntity.PUBLISHWAY, entity1.getPublishway());
        entity.setProperty(RSSEntity.CREATED, entity1.getCreated());
        entity.setProperty(RSSEntity.UPDATED, entity1.getUpdated());
        entity.setProperty(RSSEntity.LATESTDATE, entity1.getLatestdate());
        entity.setProperty(RSSEntity.DURATION, entity1.getDuration());

        datastore.put(entity);
    }

    @Override
    public void deleteEntity(Long id) {
        Key key = KeyFactory.createKey(KIND, id);
        datastore.delete(key);
    }

    public List<RSSEntity> entitiesToTasks(Iterator<Entity> resultList) {
        List<RSSEntity> resultTasks = new ArrayList<>();
        while (resultList.hasNext()) {
            resultTasks.add(entityToTask(resultList.next()));
        }
        return resultTasks;
    }

    @Override
    public List<RSSEntity> listEntities(String startCursorString) {
        FetchOptions fetchOptions = FetchOptions.Builder.withLimit(10);
        if (startCursorString != null && !startCursorString.equals("")) {
            fetchOptions.startCursor(Cursor.fromWebSafeString(startCursorString));
        }
        Query query = new Query(KIND)
                .addSort(RSSEntity.CREATED, Query.SortDirection.ASCENDING);
        PreparedQuery preparedQuery = datastore.prepare(query);
        QueryResultIterator<Entity> results = preparedQuery.asQueryResultIterator(fetchOptions);

        List<RSSEntity> resultBooks = entitiesToTasks(results);
        Cursor cursor = results.getCursor();
        if (cursor != null && resultBooks.size() == 10) {
            String cursorString = cursor.toWebSafeString();
            return resultBooks;
        } else {
            return resultBooks;
        }
    }
}
