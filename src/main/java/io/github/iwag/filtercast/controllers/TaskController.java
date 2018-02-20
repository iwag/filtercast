package io.github.iwag.filtercast.controllers;

import io.github.iwag.filtercast.models.RSSEntity;
import io.github.iwag.filtercast.repositories.TaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.*;

import java.util.*;

@RestController
public class TaskController extends BaseController {

    @Autowired
    TaskService datastoreService;

    @CrossOrigin
    @RequestMapping(path = "/rss", method = RequestMethod.GET, produces = MediaType.APPLICATION_JSON_VALUE)
    public List<RSSEntity> gets() {
        return datastoreService.listTasks(null);
    }

    @CrossOrigin
    @RequestMapping(method = RequestMethod.GET, path = "/entity/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
    public RSSEntity get(@PathVariable(name = "id", required = true) String id) {
        return datastoreService.readEntity(Long.valueOf(id));
    }

    @CrossOrigin
    @RequestMapping(path = "/rss", method = RequestMethod.PUT, produces = MediaType.APPLICATION_JSON_VALUE, consumes = MediaType.APPLICATION_JSON_VALUE)
    public RSSEntity create(@RequestBody(required = true) RSSEntity entity) {
        logger.info("RSSEntity: " + entity);

        Optional<Long> id = datastoreService.createTask(entity);
        id.ifPresent(i -> entity.setId(i.toString()));

        return entity;
    }

    @CrossOrigin
    @RequestMapping(method = RequestMethod.DELETE, path = "/entity/{id}", produces = MediaType.APPLICATION_JSON_VALUE)
    public void delete(@PathVariable(name = "id", required = true) String id) {
        datastoreService.deleteTask(Long.valueOf(id));
    }

    @CrossOrigin
    @RequestMapping(path = "/rss", method = RequestMethod.POST, produces = MediaType.APPLICATION_JSON_VALUE, consumes = MediaType.APPLICATION_JSON_VALUE)
    public void update(@RequestBody(required = true) RSSEntity entity) {
        logger.info("RSSEntity: " + entity);

        datastoreService.updateEntity(entity);

        return;
    }

}
