FROM golang:1.27-trixie AS build
WORKDIR /src

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags='-s -w' -o /out/server .

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /out/server /app/server
USER nonroot:nonroot
EXPOSE 8080
CMD ["/app/server"]
