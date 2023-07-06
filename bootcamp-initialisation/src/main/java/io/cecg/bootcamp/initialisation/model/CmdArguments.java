package io.cecg.bootcamp.initialisation.model;

import java.util.HashSet;
import java.util.Set;

public class CmdArguments {
    private final String bootcampeeRepo;
    private final String org;
    private final String gitToken;
    private final Set<String> modules;
    private final boolean helpRequest;
    private final boolean modulesRequest;

    private CmdArguments(Builder builder) {
        this.bootcampeeRepo = builder.bootcampeeRepo;
        this.org = builder.org;
        this.gitToken = builder.gitToken;
        this.modules = builder.modules;
        this.helpRequest = builder.helpRequest;
        this.modulesRequest = builder.modulesRequest;
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

    public boolean isHelpRequest() {
        return helpRequest;
    }

    public boolean isModulesRequest() {
        return modulesRequest;
    }

    public static class Builder {
        private String bootcampeeRepo;
        private String org;
        private String gitToken;
        private Set<String> modules;
        private boolean helpRequest;
        private boolean modulesRequest;

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

        public Builder setModulesRequest(boolean modulesRequest) {
            this.modulesRequest = modulesRequest;
            return this;
        }

        public Builder setModules(Set<String> modules) {
            this.modules = modules;
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