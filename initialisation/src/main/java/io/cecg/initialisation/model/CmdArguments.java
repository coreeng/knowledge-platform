package io.cecg.initialisation.model;

import java.util.HashSet;
import java.util.Set;

public class CmdArguments {
    private final String bootcampeeRepo;
    private final String org;
    private final String gitToken;
    private final Set<String> modules;
    private final String moduleLocation;
    private final boolean helpRequest;

    private CmdArguments(Builder builder) {
        this.bootcampeeRepo = builder.bootcampeeRepo;
        this.org = builder.org;
        this.gitToken = builder.gitToken;
        this.modules = builder.modules;
        this.helpRequest = builder.helpRequest;
        this.moduleLocation = builder.moduleLocation;
    }

    public String getBootcampeeRepo() {
        return bootcampeeRepo;
    }

    public String getOrg() {
        return org;
    }

    public String getGitToken() {
        return gitToken;
    }

    public Set<String> getModules() {
        return new HashSet<>(modules);
    }

    public String getModuleLocation() {
        return moduleLocation;
    }

    public boolean isHelpRequest() {
        return helpRequest;
    }

    public static class Builder {
        private String bootcampeeRepo;
        private String org;
        private String gitToken;
        private Set<String> modules;
        private String moduleLocation;
        private boolean helpRequest;

        public Builder() {
            this.modules = new HashSet<>();
        }

        public Builder setBootcampeeRepo(String bootcampeeRepo) {
            this.bootcampeeRepo = bootcampeeRepo;
            return this;
        }

        public Builder setOrg(String org) {
            this.org = org;
            return this;
        }

        public Builder setGitToken(String gitToken) {
            this.gitToken = gitToken;
            return this;
        }

        public Builder setModules(Set<String> modules) {
            this.modules = modules;
            return this;
        }

        public Builder setModuleLocation(String moduleLocation) {
            this.moduleLocation = moduleLocation;
            return this;
        }

        public Builder setHelpRequest(boolean helpRequest) {
            this.helpRequest = helpRequest;
            return this;
        }

        public CmdArguments build() {
            return new CmdArguments(this);
        }
    }
}