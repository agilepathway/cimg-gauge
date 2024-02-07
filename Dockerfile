FROM cimg/go:1.22-node

LABEL maintainer="John Boyes <john@agilepathway.co.uk>"

COPY scripts /usr/local/bin/

USER root
RUN    install-chrome.sh \
    && install-gauge.sh
USER circleci
RUN    install-gauge-plugins.sh \
    && install-taiko.sh \
    && install-mage.sh
