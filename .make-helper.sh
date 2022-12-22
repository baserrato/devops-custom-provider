#!/bin/bash

function ac () {
  printf "<         >\r" && \
  err1=$(terraform -chdir=examples/allCombined init -plugin-dir=../../.plugin-cache | grep "Error") && \
  if [[ $err1 != "" ]]; then printf "${err1}\n"; \
    else printf "<#####    >\r" && err2=$(terraform -chdir=examples/allCombined plan | grep -c "Error") && \
    if [[ $err2 -ne 0 ]]; then printf "${err2}\n"; \
      else printf "<#########>\r\n"; \
    fi; \
  fi
}

if [[ $1 = "ac" ]]; then ac; fi
