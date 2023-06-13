```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class walletFacade {
  -account : account
  -wallet : wallet
  -securityCode : securityCode
  -notification : notification
  -ledger : ledger
  +addMoneyToWallet(accountID : string, securityCode : int, amount : int) : error
  +deductMoneyFromWallet(accountID : string, securityCode : int, amount : int) : error
}

class account {
  -name : string
  +checkAccount(accountName : string) : error
}

class wallet {
  -balance : int
  +creditBalance(amount : int)
  +debitBalance(amount : int) : error
}

class securityCode {
  -code : int
  +checkCode(incomingCode : int) : error
}

class notification {
  +sendWalletCreditNotification()
  +sendWalletDebitNotification()
}

class ledger {
  +makeEntry(accountID : string, txnType : string, amount : int)
}

walletFacade --> account
walletFacade --> wallet
walletFacade --> securityCode
walletFacade --> notification
walletFacade --> ledger
```
```plantuml
left to right direction
skinparam backgroundColor #F0F0F0

class Projector {
  +On()
  +Off()
}

class DVDPlayer {
  +Play()
  +Stop()
}

class SoundSystem {
  +On()
  +Off()
}

class HomeTheaterFacade {
  +WatchMovie()
  +StopMovie()
}

Projector "1" -down- "1" HomeTheaterFacade
DVDPlayer "1" -- "1" HomeTheaterFacade
SoundSystem "1" -- "1" HomeTheaterFacade
```