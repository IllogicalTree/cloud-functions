# Image resize service

A dead simple image resizing service designed to be deployed to a google cloud function.

## Usage

The service can be deployed to google cloud functions with the following command assuming you are authenticated and have cloud functions enabled


```
gcloud functions deploy ImageResize --runtime go113 --trigger-http --allow-unauthenticated
```

Running locally is development mode is as simple as running the following command; however you will want to rename the package name to main when running this way, (A bug with deploying to gcloud functions meant that a package named main fails to deploy).

```
go run .
```

## Resources

https://cloud.google.com/functions/docs/quickstart-go
https://github.com/nfnt/resize

## Notes

Performance is not great when resizing larger images and may cause the cloud function to time out before the image transformation completes. If I come ever come back to this project in the future it might be worth using a c library with go bindings for better performance or using a project like [this](https://github.com/h2non/imaginary).
