```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class Profile {
  -Email : string
}

interface ProfileIterator {
  +GetNext() : Profile
  +HasMore() : bool
}

interface SocialNetwork {
  +CreateFriendsIterator(profileId: string) : ProfileIterator
  +CreateCoworkersIterator(profileId: string) : ProfileIterator
}

class WeChat {
  +CreateFriendsIterator(profileId: string) : ProfileIterator
  +CreateCoworkersIterator(profileId: string) : ProfileIterator
}

class WeChatIterator {
  -weChat : WeChat
  -profileId : string
  -typeStr : string
  -position : int
  -cache : Profile[]
  +GetNext() : Profile
  +HasMore() : bool
}

class SocialSpammer {
  +Send(iterator : ProfileIterator, message: string)
}

class Application {
  -network : SocialNetwork
  -spammer : SocialSpammer
  +Config()
  +SendSpamToFriends(profileId: string)
  +SendSpamToCoworkers(profileId: string)
}

WeChatIterator ..> WeChat
WeChatIterator ..> Profile
WeChat --|> SocialNetwork
WeChatIterator --|> ProfileIterator
Application o-- SocialNetwork
Application o-- SocialSpammer
SocialSpammer ..> ProfileIterator
```