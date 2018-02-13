package io.github.iwag.filtercast.models;

public class RSSEntity {
    public static final String ID = "__id__";
    public static final String URL = "__url__";
    public static final String HISTORY = "__history__";
    public static final String PUBLISHWAY = "publishway";
    public static final String CREATED = "created";
    public static final String UPDATED = "updated";
    public static final String LATESTDATE = "latest";
    public static final String DURATION = "duration";

    private String id;
    private String url;
    private String history;
    private String publishway;
    private String created;
    private String updated;
    private String latestdate;
    private String duration;

    public RSSEntity() {

    }

    public RSSEntity(String id, String url, String history, String publishway, String created, String updated, String latestdate, String duration) {
        this.id = id;
        this.url = url;
        this.history = history;
        this.publishway = publishway;
        this.created = created;
        this.updated = updated;
        this.latestdate = latestdate;
        this.duration = duration;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getHistory() {
        return history;
    }

    public void setHistory(String history) {
        this.history = history;
    }

    public String getPublishway() {
        return publishway;
    }

    public void setPublishway(String publishway) {
        this.publishway = publishway;
    }

    public String getCreated() {
        return created;
    }

    public void setCreated(String created) {
        this.created = created;
    }

    public String getUpdated() {
        return updated;
    }

    public void setUpdated(String updated) {
        this.updated = updated;
    }

    public String getLatestdate() {
        return latestdate;
    }

    public void setLatestdate(String latestdate) {
        this.latestdate = latestdate;
    }

    public String getDuration() {
        return duration;
    }

    public void setDuration(String duration) {
        this.duration = duration;
    }

    @Override
    public String toString() {
        return "RSSEntity{" +
                "id='" + id + '\'' +
                ", url='" + url + '\'' +
                ", history='" + history + '\'' +
                ", publishway='" + publishway + '\'' +
                ", created='" + created + '\'' +
                ", updated='" + updated + '\'' +
                ", latestdate='" + latestdate + '\'' +
                ", duration='" + duration + '\'' +
                '}';
    }

    public RSSEntity(Builder builder) {
    }

    public static class Builder {
        private Long id;
        private String url;
        private String history;
        private String publishway;
        private String created;
        private String updated;
        private String latestdate;
        private String duration;

        public Builder url(String url) {
            this.url = url;
            return this;
        }

        public Builder history(String history) {
            this.history = history;
            return this;
        }

        public Builder publishway(String publishway) {
            this.publishway = publishway;
            return this;
        }


        public Builder created(String created) {
            this.created = created;
            return this;
        }

        public Builder updated(String updated) {
            this.updated = updated;
            return this;
        }

        public Builder latestdate(String latestdate) {
            this.latestdate = latestdate;
            return this;
        }

        public Builder duration(String duration) {
            this.duration = duration;
            return this;
        }

        public Builder id(Long id) {
            this.id = id;
            return this;
        }

        public RSSEntity build() {
            return new RSSEntity(this);
        }
    }
}
