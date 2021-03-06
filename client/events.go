package client

//go:generate stringer -type=EventType
type EventType uint16

const (
	evUnused                               EventType = 0
	evLeave                                EventType = 1
	evJoinFinished                         EventType = 2
	evMove                                 EventType = 3
	evTeleport                             EventType = 4
	evChangeEquipment                      EventType = 5
	evHealthUpdate                         EventType = 6
	evEnergyUpdate                         EventType = 7
	evDamageShieldUpdate                   EventType = 8
	evCraftingFocusUpdate                  EventType = 9
	evActiveSpellEffectsUpdate             EventType = 10
	evResetCooldowns                       EventType = 11
	evAttack                               EventType = 12
	evCastStart                            EventType = 13
	evCastCancel                           EventType = 14
	evCastTimeUpdate                       EventType = 15
	evCastFinished                         EventType = 16
	evCastSpell                            EventType = 17
	evCastHit                              EventType = 18
	evCastHits                             EventType = 19
	evChannelingEnded                      EventType = 20
	evAttackBuilding                       EventType = 21
	evInventoryPutItem                     EventType = 22
	evInventoryDeleteItem                  EventType = 23
	evNewCharacter                         EventType = 24
	evNewEquipmentItem                     EventType = 25
	evNewSimpleItem                        EventType = 26
	evNewFurnitureItem                     EventType = 27
	evNewJournalItem                       EventType = 28
	evNewLaborerItem                       EventType = 29
	evNewSimpleHarvestableObject           EventType = 30
	evNewSimpleHarvestableObjectList       EventType = 31
	evNewHarvestableObject                 EventType = 32
	evNewSilverObject                      EventType = 33
	evNewBuilding                          EventType = 34
	evHarvestableChangeState               EventType = 35
	evMobChangeState                       EventType = 36
	evFactionBuildingInfo                  EventType = 37
	evCraftBuildingInfo                    EventType = 38
	evRepairBuildingInfo                   EventType = 39
	evMeldBuildingInfo                     EventType = 40
	evConstructionSiteInfo                 EventType = 41
	evPlayerBuildingInfo                   EventType = 42
	evFarmBuildingInfo                     EventType = 43
	evTutorialBuildingInfo                 EventType = 44
	evLaborerObjectInfo                    EventType = 45
	evLaborerObjectJobInfo                 EventType = 46
	evMarketPlaceBuildingInfo              EventType = 47
	evHarvestStart                         EventType = 48
	evHarvestCancel                        EventType = 49
	evHarvestFinished                      EventType = 50
	evTakeSilver                           EventType = 51
	evActionOnBuildingStart                EventType = 52
	evActionOnBuildingCancel               EventType = 53
	evActionOnBuildingFinished             EventType = 54
	evItemRerollQualityStart               EventType = 55
	evItemRerollQualityCancel              EventType = 56
	evItemRerollQualityFinished            EventType = 57
	evInstallResourceStart                 EventType = 58
	evInstallResourceCancel                EventType = 59
	evInstallResourceFinished              EventType = 60
	evCraftItemFinished                    EventType = 61
	evLogoutCancel                         EventType = 62
	evChatMessage                          EventType = 63
	evChatSay                              EventType = 64
	evChatWhisper                          EventType = 65
	evChatMuted                            EventType = 66
	evPlayEmote                            EventType = 67
	evStopEmote                            EventType = 68
	evSystemMessage                        EventType = 69
	evUtilityTextMessage                   EventType = 70
	evUpdateMoney                          EventType = 71
	evUpdateFame                           EventType = 72
	evUpdateLearningPoints                 EventType = 73
	evUpdateReSpecPoints                   EventType = 74
	evUpdateCurrency                       EventType = 75
	evUpdateFactionStanding                EventType = 76
	evRespawn                              EventType = 77
	evServerDebugLog                       EventType = 78
	evCharacterEquipmentChanged            EventType = 79
	evRegenerationHealthChanged            EventType = 80
	evRegenerationEnergyChanged            EventType = 81
	evRegenerationMountHealthChanged       EventType = 82
	evRegenerationCraftingChanged          EventType = 83
	evRegenerationHealthEnergyComboChanged EventType = 84
	evRegenerationPlayerComboChanged       EventType = 85
	evDurabilityChanged                    EventType = 86
	evNewLoot                              EventType = 87
	evAttachItemContainer                  EventType = 88
	evDetachItemContainer                  EventType = 89
	evInvalidateItemContainer              EventType = 90
	evLockItemContainer                    EventType = 91
	evGuildUpdate                          EventType = 92
	evGuildPlayerUpdated                   EventType = 93
	evInvitedToGuild                       EventType = 94
	evGuildMemberWorldUpdate               EventType = 95
	evUpdateMatchDetails                   EventType = 96
	evObjectEvent                          EventType = 97
	evNewMonolithObject                    EventType = 98
	evNewSiegeCampObject                   EventType = 99
	evNewOrbObject                         EventType = 100
	evNewCastleObject                      EventType = 101
	evNewSpellEffectArea                   EventType = 102
	evNewChainSpell                        EventType = 103
	evUpdateChainSpell                     EventType = 104
	evNewTreasureChest                     EventType = 105
	evStartMatch                           EventType = 106
	evStartTerritoryMatchInfos             EventType = 107
	evStartArenaMatchInfos                 EventType = 108
	evEndTerritoryMatch                    EventType = 109
	evEndArenaMatch                        EventType = 110
	evMatchUpdate                          EventType = 111
	evActiveMatchUpdate                    EventType = 112
	evNewMob                               EventType = 113
	evDebugAggroInfo                       EventType = 114
	evDebugVariablesInfo                   EventType = 115
	evDebugReputationInfo                  EventType = 116
	evDebugDiminishingReturnInfo           EventType = 117
	evDebugSmartClusterQueueInfo           EventType = 118
	evClaimOrbStart                        EventType = 119
	evClaimOrbFinished                     EventType = 120
	evClaimOrbCancel                       EventType = 121
	evOrbUpdate                            EventType = 122
	evOrbClaimed                           EventType = 123
	evNewWarCampObject                     EventType = 124
	evNewMatchLootChestObject              EventType = 125
	evNewArenaExit                         EventType = 126
	evGuildMemberTerritoryUpdate           EventType = 127
	evInvitedMercenaryToMatch              EventType = 128
	evClusterInfoUpdate                    EventType = 129
	evForcedMovement                       EventType = 130
	evForcedMovementCancel                 EventType = 131
	evCharacterStats                       EventType = 132
	evCharacterStatsKillHistory            EventType = 133
	evCharacterStatsDeathHistory           EventType = 134
	evGuildStats                           EventType = 135
	evKillHistoryDetails                   EventType = 136
	evFullAchievementInfo                  EventType = 137
	evFinishedAchievement                  EventType = 138
	evAchievementProgressInfo              EventType = 139
	evFullAchievementProgressInfo          EventType = 140
	evFullTrackedAchievementInfo           EventType = 141
	evFullAutoLearnAchievementInfo         EventType = 142
	evQuestGiverQuestOffered               EventType = 143
	evQuestGiverDebugInfo                  EventType = 144
	evConsoleEvent                         EventType = 145
	evTimeSync                             EventType = 146
	evChangeAvatar                         EventType = 147
	evChangeMountSkin                      EventType = 148
	evGameEvent                            EventType = 149
	evKilledPlayer                         EventType = 150
	evDied                                 EventType = 151
	evKnockedDown                          EventType = 152
	evMatchPlayerJoinedEvent               EventType = 153
	evMatchPlayerStatsEvent                EventType = 154
	evMatchPlayerStatsCompleteEvent        EventType = 155
	evMatchTimeLineEventEvent              EventType = 156
	evMatchPlayerMainGearStatsEvent        EventType = 157
	evMatchPlayerChangedAvatarEvent        EventType = 158
	evInvitationPlayerTrade                EventType = 159
	evPlayerTradeStart                     EventType = 160
	evPlayerTradeCancel                    EventType = 161
	evPlayerTradeUpdate                    EventType = 162
	evPlayerTradeFinished                  EventType = 163
	evPlayerTradeAcceptChange              EventType = 164
	evMiniMapPing                          EventType = 165
	evMarketPlaceNotification              EventType = 166
	evDuellingChallengePlayer              EventType = 167
	evNewDuellingPost                      EventType = 168
	evDuelStarted                          EventType = 169
	evDuelEnded                            EventType = 170
	evDuelDenied                           EventType = 171
	evDuelLeftArea                         EventType = 172
	evDuelReEnteredArea                    EventType = 173
	evNewRealEstate                        EventType = 174
	evMiniMapOwnedBuildingsPositions       EventType = 175
	evRealEstateListUpdate                 EventType = 176
	evGuildLogoUpdate                      EventType = 177
	evGuildLogoChanged                     EventType = 178
	evPlaceableObjectPlace                 EventType = 179
	evPlaceableObjectPlaceCancel           EventType = 180
	evFurnitureObjectBuffProviderInfo      EventType = 181
	evFurnitureObjectCheatProviderInfo     EventType = 182
	evFarmableObjectInfo                   EventType = 183
	evNewUnreadMails                       EventType = 184
	evGuildLogoObjectUpdate                EventType = 185
	evStartLogout                          EventType = 186
	evNewChatChannels                      EventType = 187
	evJoinedChatChannel                    EventType = 188
	evLeftChatChannel                      EventType = 189
	evRemovedChatChannel                   EventType = 190
	evAccessStatus                         EventType = 191
	evMounted                              EventType = 192
	evMountStart                           EventType = 193
	evMountCancel                          EventType = 194
	evNewTravelpoint                       EventType = 195
	evNewIslandAccessPoint                 EventType = 196
	evNewExit                              EventType = 197
	evUpdateHome                           EventType = 198
	evUpdateChatSettings                   EventType = 199
	evResurrectionOffer                    EventType = 200
	evResurrectionReply                    EventType = 201
	evLootEquipmentChanged                 EventType = 202
	evUpdateUnlockedGuildLogos             EventType = 203
	evUpdateUnlockedAvatars                EventType = 204
	evUpdateUnlockedAvatarRings            EventType = 205
	evUpdateUnlockedBuildings              EventType = 206
	evNewIslandManagement                  EventType = 207
	evNewTeleportStone                     EventType = 208
	evCloak                                EventType = 209
	evPartyInvitation                      EventType = 210
	evPartyJoined                          EventType = 211
	evPartyDisbanded                       EventType = 212
	evPartyPlayerJoined                    EventType = 213
	evPartyChangedOrder                    EventType = 214
	evPartyPlayerLeft                      EventType = 215
	evPartyLeaderChanged                   EventType = 216
	evPartyLootSettingChangedPlayer        EventType = 217
	evPartySilverGained                    EventType = 218
	evPartyPlayerUpdated                   EventType = 219
	evPartyInvitationPlayerBusy            EventType = 220
	evPartyMarkedObjectsUpdated            EventType = 221
	evPartyOnClusterPartyJoined            EventType = 222
	evPartySetRoleFlag                     EventType = 223
	evSpellCooldownUpdate                  EventType = 224
	evNewHellgate                          EventType = 225
	evNewHellgateExit                      EventType = 226
	evNewExpeditionExit                    EventType = 227
	evNewExpeditionNarrator                EventType = 228
	evExitEnterStart                       EventType = 229
	evExitEnterCancel                      EventType = 230
	evExitEnterFinished                    EventType = 231
	evHellClusterTimeUpdate                EventType = 232
	evNewQuestGiverObject                  EventType = 233
	evFullQuestInfo                        EventType = 234
	evQuestProgressInfo                    EventType = 235
	evQuestGiverInfoForPlayer              EventType = 236
	evFullExpeditionInfo                   EventType = 237
	evExpeditionQuestProgressInfo          EventType = 238
	evInvitedToExpedition                  EventType = 239
	evExpeditionRegistrationInfo           EventType = 240
	evEnteringExpeditionStart              EventType = 241
	evEnteringExpeditionCancel             EventType = 242
	evRewardGranted                        EventType = 243
	evArenaRegistrationInfo                EventType = 244
	evEnteringArenaStart                   EventType = 245
	evEnteringArenaCancel                  EventType = 246
	evEnteringArenaLockStart               EventType = 247
	evEnteringArenaLockCancel              EventType = 248
	evInvitedToArenaMatch                  EventType = 249
	evPlayerCounts                         EventType = 250
	evInCombatStateUpdate                  EventType = 251
	evOtherGrabbedLoot                     EventType = 252
	evSiegeCampClaimStart                  EventType = 253
	evSiegeCampClaimCancel                 EventType = 254
	evSiegeCampClaimFinished               EventType = 255
	evSiegeCampScheduleResult              EventType = 256
	evTreasureChestUsingStart              EventType = 257
	evTreasureChestUsingFinished           EventType = 258
	evTreasureChestUsingCancel             EventType = 259
	evTreasureChestUsingOpeningComplete    EventType = 260
	evTreasureChestForceCloseInventory     EventType = 261
	evPremiumChanged                       EventType = 262
	evPremiumExtended                      EventType = 263
	evPremiumLifeTimeRewardGained          EventType = 264
	evLaborerGotUpgraded                   EventType = 265
	evJournalGotFull                       EventType = 266
	evJournalFillError                     EventType = 267
	evFriendRequest                        EventType = 268
	evFriendRequestInfos                   EventType = 269
	evFriendInfos                          EventType = 270
	evFriendRequestAnswered                EventType = 271
	evFriendOnlineStatus                   EventType = 272
	evFriendRequestCanceled                EventType = 273
	evFriendRemoved                        EventType = 274
	evFriendUpdated                        EventType = 275
	evPartyLootItems                       EventType = 276
	evPartyLootItemsRemoved                EventType = 277
	evReputationUpdate                     EventType = 278
	evDefenseUnitAttackBegin               EventType = 279
	evDefenseUnitAttackEnd                 EventType = 280
	evDefenseUnitAttackDamage              EventType = 281
	evUnrestrictedPvpZoneUpdate            EventType = 282
	evReputationImplicationUpdate          EventType = 283
	evNewMountObject                       EventType = 284
	evMountHealthUpdate                    EventType = 285
	evMountCooldownUpdate                  EventType = 286
	evNewExpeditionAgent                   EventType = 287
	evNewExpeditionCheckPoint              EventType = 288
	evExpeditionStartEvent                 EventType = 289
	evVoteEvent                            EventType = 290
	evRatingEvent                          EventType = 291
	evNewArenaAgent                        EventType = 292
	evBoostFarmable                        EventType = 293
	evUseFunction                          EventType = 294
	evNewPortalEntrance                    EventType = 295
	evNewPortalExit                        EventType = 296
	evNewRandomDungeonExit                 EventType = 297
	evWaitingQueueUpdate                   EventType = 298
	evPlayerMovementRateUpdate             EventType = 299
	evObserveStart                         EventType = 300
	evMinimapZergs                         EventType = 301
	evPaymentTransactions                  EventType = 302
	evPerformanceStatsUpdate               EventType = 303
	evOverloadModeUpdate                   EventType = 304
	evDebugDrawEvent                       EventType = 305
	evRecordCameraMove                     EventType = 306
	evRecordStart                          EventType = 307
	evTerritoryClaimStart                  EventType = 308
	evTerritoryClaimCancel                 EventType = 309
	evTerritoryClaimFinished               EventType = 310
	evTerritoryScheduleResult              EventType = 311
	evUpdateAccountState                   EventType = 312
	evStartDeterministicRoam               EventType = 313
	evGuildFullAccessTagsUpdated           EventType = 314
	evGuildAccessTagUpdated                EventType = 315
	evGvgSeasonUpdate                      EventType = 316
	evGvgSeasonCheatCommand                EventType = 317
	evSeasonPointsByKillingBooster         EventType = 318
	evFishingStart                         EventType = 319
	evFishingCast                          EventType = 320
	evFishingCatch                         EventType = 321
	evFishingFinished                      EventType = 322
	evFishingCancel                        EventType = 323
	evNewFloatObject                       EventType = 324
	evNewFishingZoneObject                 EventType = 325
	evFishingMiniGame                      EventType = 326
	evSteamAchievementCompleted            EventType = 327
	evUpdatePuppet                         EventType = 328
	evChangeFlaggingFinished               EventType = 329
	evNewOutpostObject                     EventType = 330
	evOutpostUpdate                        EventType = 331
	evOutpostClaimed                       EventType = 332
	evOutpostReward                        EventType = 333
	evOverChargeEnd                        EventType = 334
	evOverChargeStatus                     EventType = 335
	evPartyFinderFullUpdate                EventType = 336
	evPartyFinderUpdate                    EventType = 337
	evPartyFinderApplicantsUpdate          EventType = 338
	evPartyFinderEquipmentSnapshot         EventType = 339
	evPartyFinderJoinRequestDeclined       EventType = 340
	evNewUnlockedPersonalSeasonRewards     EventType = 341
	evPersonalSeasonPointsGained           EventType = 342
	evEasyAntiCheatMessageToClient         EventType = 343
	evMatchLootChestOpeningStart           EventType = 344
	evMatchLootChestOpeningFinished        EventType = 345
	evMatchLootChestOpeningCancel          EventType = 346
	evNotifyCrystalMatchReward             EventType = 347
	evCrystalRealmFeedback                 EventType = 348
	evNewLocationMarker                    EventType = 349
	evNewTutorialBlocker                   EventType = 350
	evNewTileSwitch                        EventType = 351
	evNewInformationProvider               EventType = 352
	evNewDynamicGuildLogo                  EventType = 353
	evTutorialUpdate                       EventType = 354
	evTriggerHintBox                       EventType = 355
	evRandomDungeonPositionInfo            EventType = 356
	evNewLootChest                         EventType = 357
	evUpdateLootChest                      EventType = 358
	evLootChestOpened                      EventType = 359
	evNewShrine                            EventType = 360
	evUpdateShrine                         EventType = 361
	evMutePlayerUpdate                     EventType = 362
	evShopTileUpdate                       EventType = 363
	evShopUpdate                           EventType = 364
	evEasyAntiCheatKick                    EventType = 365
	evUnlockVanityUnlock                   EventType = 366
	evAvatarUnlocked                       EventType = 367
	evCustomizationChanged                 EventType = 368
	evBaseVaultInfo                        EventType = 369
	evGuildVaultInfo                       EventType = 370
	evBankVaultInfo                        EventType = 371
	evRecoveryVaultPlayerInfo              EventType = 372
	evRecoveryVaultGuildInfo               EventType = 373
	evUpdateWardrobe                       EventType = 374
	evCastlePhaseChanged                   EventType = 375
	evGuildAccountLogEvent                 EventType = 376
	evNewHideoutObject                     EventType = 377
	evNewHideoutManagement                 EventType = 378
	evNewHideoutExit                       EventType = 379
	evInitHideoutAttackStart               EventType = 380
	evInitHideoutAttackCancel              EventType = 381
	evInitHideoutAttackFinished            EventType = 382
	evHideoutManagementUpdate              EventType = 383
	evIpChanged                            EventType = 384
	evSmartClusterQueueUpdateInfo          EventType = 385
	evSmartClusterQueueActiveInfo          EventType = 386
	evSmartClusterQueueKickWarning         EventType = 387
	evSmartClusterQueueInvite              EventType = 388
	evReceivedGvgSeasonPoints              EventType = 389
	evTerritoryBonusLevelUpdate            EventType = 390
	evOpenWorldAttackScheduleStart         EventType = 391
	evOpenWorldAttackScheduleFinished      EventType = 392
	evOpenWorldAttackScheduleCancel        EventType = 393
	evOpenWorldAttackConquerStart          EventType = 394
	evOpenWorldAttackConquerFinished       EventType = 395
	evOpenWorldAttackConquerCancel         EventType = 396
	evOpenWorldAttackConquerStatus         EventType = 397
	evOpenWorldAttackStart                 EventType = 398
	evOpenWorldAttackEnd                   EventType = 399
	evNewRandomResourceBlocker             EventType = 400
	evNewHomeObject                        EventType = 401
	evHideoutObjectUpdate                  EventType = 402
)
