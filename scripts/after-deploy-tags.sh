#!/bin/bash

FULL_CHROME_VERSION=$(docker run --rm -i "$IMAGE_NAME" bash -c 'echo "$(google-chrome --version)"')
CHROME_VERSION=${FULL_CHROME_VERSION//[!0-9.]/}
CHROME_TAG="CHROME-${CHROME_VERSION}"

CIRCLECI_TAG="CIRCLECI-${CIRCLE_BUILD_NUM}"

GAUGE_VERSION=$(docker run --rm -i "$IMAGE_NAME" bash -c 'echo "$(gauge-version)"')
GAUGE_TAG="GAUGE-$GAUGE_VERSION"

NODE_VERSION=13.12
NODE_TAG="NODE-${NODE_VERSION}"

# Use Gauge version as the semantic version of the image, as this is a Gauge image and
# CircleCI use the semantic version of the main tool or language that is being provided
# for their base images: https://circleci.com/docs/2.0/circleci-images/#best-practices
# Also tag the Gauge version and the circle build together so that consumers can pin to an 
# idempotent image
echo "$GAUGE_VERSION","$GAUGE_VERSION"-"$CIRCLECI_TAG","$GAUGE_TAG","$CHROME_TAG","$CIRCLECI_TAG","$NODE_TAG"