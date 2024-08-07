FROM arm64v8/golang:1.22 as build
WORKDIR /build

ARG CGO_ENABLED=0
ENV CGO_ENABLED=${CGO_ENABLED}

ARG GOOS=linux
ARG GOOS=${GOOS}

ARG GO111MODULE=on
ENV GO111MODULE=${GO111MODULE}

COPY go.mod go.sum ./
COPY . .

RUN go build -tags lambda.norpc -o image-upload ./cmd/lambdas/images/upload/main.go && \
    go build -tags lambda.norpc -o image-resize ./cmd/lambdas/images/resize/main.go && \
    go build -tags lambda.norpc -o blog-update ./cmd/lambdas/blog/update/main.go && \
    go build -tags lambda.norpc -o blog-notify ./cmd/lambdas/blog/notify/main.go

FROM public.ecr.aws/lambda/provided:al2023 AS image-upload-lambda
COPY --from=build /build/image-upload ./image-upload
ENTRYPOINT [ "./image-upload" ]

FROM public.ecr.aws/lambda/provided:al2023 AS image-resize-lambda
COPY --from=build /build/image-resize ./image-resize
ENTRYPOINT [ "./image-resize" ]

FROM public.ecr.aws/lambda/provided:al2023 AS blog-update-lambda
COPY --from=build /build/blog-update ./blog-update
ENTRYPOINT [ "./blog-update" ]

FROM public.ecr.aws/lambda/provided:al2023 AS blog-notify-lambda
COPY --from=build /build/blog-notify ./blog-notify
ENTRYPOINT [ "./blog-notify" ]