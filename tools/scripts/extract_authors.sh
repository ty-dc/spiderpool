#!/usr/bin/env bash

# Copyright 2022 Authors of spidernet-io
# SPDX-License-Identifier: Apache-2.0

# Ensure sort order doesn't depend on locale
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8

# exclusion list
# Exclude bots and non-compliant lists
excluded_name=("dependabot[bot]" "root" "weizhou.lan@daocloud.io")

function extract_authors() {
	authors=$(git shortlog --summary \
		  | awk '{$1=""; print $0}' \
		  | sed -e 's/^ //' \
			-e '/vagrant/d')

	# Iterate $authors by line
	IFS=$'\n'
    # Exclude emails that do not meet specifications and are duplicates.
    email=()
	for i in $authors; do
        # Skip excluded names
        if [[ " ${excluded_name[@]} " =~ " ${i} " ]]; then
            continue
        fi
    	new_email=$(git log --use-mailmap --author="$i" --format="%<|(40)%aN%aE" \
        | grep -v noreply.github.com \
        | grep -v example.com \
        | head -n 1 \
        | awk '{$1=""; print $0}')
        if [[ " ${email[@]} " =~ " ${new_email// /} " ]]; then
            continue
        else
            email+=$new_email
            git log --use-mailmap --author="$i" --format="%<|(40)%aN%aE" | grep -v noreply.github.com | grep -v example.com | head -n 1
        fi
	done
}

extract_authors | sort -u