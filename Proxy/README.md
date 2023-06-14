```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

interface ThirdPartyYouTubeLib {
    + listVideos(): string
    + getVideoInfo(id: string): string
    + downloadVideo(id: string)
}

class ThirdPartyYouTubeClass {
    + listVideos(): string
    + getVideoInfo(id: string): string
    + downloadVideo(id: string)
}

class CachedYouTubeClass {
    - service: ThirdPartyYouTubeLib
    - listCache: string
    - videoCache: string
    - needReset: boolean
    + CachedYouTubeClass(service: ThirdPartyYouTubeLib)
    + listVideos(): string
    + getVideoInfo(id: string): string
    + downloadVideo(id: string)
    + downloadExists(id: string): boolean
}

class YouTubeManager {
    - service: ThirdPartyYouTubeLib
    + YouTubeManager(service: ThirdPartyYouTubeLib)
    + renderVideoPage(id: string)
    + renderListPanel()
    + reactOnUserInput()
}

class Application {
    + init()
}

ThirdPartyYouTubeLib <|.. ThirdPartyYouTubeClass
ThirdPartyYouTubeLib <|.. CachedYouTubeClass
ThirdPartyYouTubeLib <-- YouTubeManager
CachedYouTubeClass o-- ThirdPartyYouTubeLib
Application o-- YouTubeManager
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class RemoteImage {
    - url: String
    + RemoteImage(url: String)
    + displayImage(): void
}

class ImageProxy {
    - url: String
    - image: RemoteImage
    + ImageProxy(url: String)
    + displayImage(): void
}

RemoteImage --> ImageProxy
```