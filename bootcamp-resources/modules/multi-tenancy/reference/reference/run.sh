#!/bin/bash
echo "Onboarding the tenants into the cluster by creating all resources needed"
cue cmd onboard ./onboard
echo "Create monitoring resources for all the tenants that require it"
cue cmd monitoring ./onboard