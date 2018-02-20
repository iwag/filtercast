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
public class DatastoreService implements TaskService {

    private final static String KIND = "Task2";

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
    public Optional<Long> createTask(RSSEntity task) {
        Entity incTaskEntity = new Entity(KIND);
        incTaskEntity.setProperty(RSSEntity.HISTORY, task.getHistory());
        incTaskEntity.setProperty(RSSEntity.URL, task.getUrl());
        incTaskEntity.setProperty(RSSEntity.PUBLISHWAY, task.getHistory());
        incTaskEntity.setProperty(RSSEntity.CREATED, task.getCreated());
        incTaskEntity.setProperty(RSSEntity.UPDATED, task.getUpdated());
        incTaskEntity.setProperty(RSSEntity.LATESTDATE, task.getLatestdate());
        incTaskEntity.setProperty(RSSEntity.DURATION, task.getDuration());

        Key k = datastore.put(incTaskEntity);
        return Optional.of(k.getId());
    }

    @Override
    public RSSEntity readEntity(Long taskId) {
        Entity entity = null;
        try {
            entity = datastore.get(KeyFactory.createKey(KIND, taskId));
        } catch (EntityNotFoundException e) {
            return null;
        }
        return entityToTask(entity);
    }

    @Override
    public void updateEntity(RSSEntity task) {
        Key key = KeyFactory.createKey(KIND, task.getId());
        Entity entity = new Entity(key);
        entity.setProperty(RSSEntity.HISTORY, task.getHistory());
        entity.setProperty(RSSEntity.URL, task.getUrl());
        entity.setProperty(RSSEntity.PUBLISHWAY, task.getPublishway());
        entity.setProperty(RSSEntity.CREATED, task.getCreated());
        entity.setProperty(RSSEntity.UPDATED, task.getUpdated());
        entity.setProperty(RSSEntity.LATESTDATE, task.getLatestdate());
        entity.setProperty(RSSEntity.DURATION, task.getDuration());

        datastore.put(entity);
    }

    @Override
    public void deleteTask(Long taskId) {
        Key key = KeyFactory.createKey(KIND, taskId);
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
    public List<RSSEntity> listTasks(String startCursorString) {
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
