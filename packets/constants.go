package packets

type Packet struct {
	Data interface{}
	Type int32
}

type FinalPacket struct {
	Content []byte
	// Ignored is a series of users of which this packet should NEVER arrive.
	Ignored []string
}

const (
	BYTE = iota
	UINT16
	SINT16
	UINT32
	SINT32
	UINT64
	SINT64
	STRING
	FLOAT
	BYTES
	INT_LIST
)

const (
	OsuSendUserState              = iota // byte (actionID), string (map name), string (map md5), int32 (mods), byte (mode), int32 (beatmap /b/ ID) - update osu about the user state
	OsuSendIRCMessage                    // emptystring, string (content), string (channel), emptyint32 - receive message from in-game chat
	OsuExit                              // emptyint32 - osu closes
	OsuRequestStatusUpdate               // Null - update player stats
	OsuPong                              // Null - ping callback
	BanchoLoginReply                     // int32 - user ID or fail ID
	BanchoCommandError                   // reply to an error
	BanchoSendMessage                    // string (username), string (content), string (channel), int32 (userid) - Add a message to a channel
	BanchoPing                           // ping request
	BanchoHandleIRCUsernameChange        // someone changes name in irc
	BanchoHandleIRCQuit                  // someone logged out
	BanchoHandleUserUpdate               // packets.UserDataFullInfo - in-depth user info (rank, pp, level, score...)
	BanchoHandleUserQuit                 // int32, byte? - user has quit (broadcasted to all users)
	BanchoSpectatorJoined                // new spec
	BanchoSpectatorLeft                  // spectator left
	BanchoSpectateFrames                 // spectator frames chunks
	OsuStartSpectating                   // request to spectate someone
	OsuStopSpectating                    // stop spectating
	OsuSpectateFrames                    // spectator frames (client packet not from bancho unlike BanchoSpectateFrames)
	BanchoVersionUpdate                  // check for updates
	OsuErrorReport                       // report error to osu.ppy.sh
	OsuCantSpectate                      // can't spectate the host for whatever reason
	BanchoSpectatorCantSpectate          // can't spectate because no map
	BanchoGetAttention                   // make osu popup
	BanchoAnnounce                       // announcement popup
	OsuSendIRCMessagePrivate             // not sure
	BanchoMatchUpdate                    // update match details
	BanchoMatchNew                       // new match
	BanchoMatchDisband                   // close room
	OsuLobbyPart                         // Null - client left lobby
	OsuLobbyJoin                         // Null - client joined lobby
	OsuMatchCreate                       // client created a new lobby
	OsuMatchJoin                         // sends a request to bancho (join lobby)
	OsuMatchPart
	BanchoLobbyJoinOBSOLETE // according to the mid-2014 decompiled code this is when bancho informs a client about a new player that joins a lobby this is obsolete now.
	BanchoLobbyPartOBSOLETE // according to the mid-2014 decompiled code this is when bancho informs a client about a new player that joins a lobby this is obsolete now.
	BanchoMatchJoinSuccess
	BanchoMatchJoinFail
	OsuMatchChangeSlot
	OsuMatchReady
	OsuMatchLock
	OsuMatchChangeSettings
	BanchoFellowSpectatorJoined
	BanchoFellowSpectatorLeft
	OsuMatchStart
	AllPlayersLoaded // no one is missing beatmap
	BanchoMatchStart
	OsuMatchScoreUpdate
	BanchoMatchScoreUpdate
	OsuMatchComplete
	BanchoMatchTransferHost
	OsuMatchChangeMods
	OsuMatchLoadComplete
	BanchoMatchAllPlayersLoaded
	OsuMatchNoBeatmap
	OsuMatchNotReady
	OsuMatchFailed
	BanchoMatchPlayerFailed
	BanchoMatchComplete
	OsuMatchHasBeatmap
	OsuMatchSkipRequest
	BanchoMatchSkip
	BanchoUnauthorised
	OsuChannelJoin           // string - not hard to guess what it is and what it sends
	BanchoChannelJoinSuccess // string - Tells the client they have been successfully subscribed to a channel.
	BanchoChannelAvailable   // string, string, short - Channel name, description and current number of users.
	BanchoChannelRevoked     // string - Channel to remove from the client.
	BanchoChannelAvailableAutojoin
	OsuBeatmapInfoRequest // []string? Looks like array length (uint32 this time) and then lotsa strings. Requests info about beatmaps, requiring a subsequent response with BanchoBeatmapInfoReply I suppose.
	BanchoBeatmapInfoReply
	OsuMatchTransferHost
	BanchoLoginPermissions // int32 - See packets.UserPrivileges constants
	BanchoFriendList       // []int32 - ALL FRIENDS, not just the ones online.
	OsuFriendAdd           // int32 - user id with the friend to add
	OsuFriendRemove        // int32 - user id with the friend to delete
	BanchoProtocolVersion  // int32 - bancho protocol version (always 19 for the moment)
	BanchoTitleUpdate
	OsuMatchChangeTeam
	OsuChannelLeave   // string - channel to part
	OsuReceiveUpdates // emptyint32 - ?
	BanchoMonitor
	BanchoMatchPlayerSkipped
	OsuSetIrcAwayMessage // This packet has too many memes. Will implement later on.
	BanchoUserPresence   // int32, string, byte, byte, byte, float, float, int - Basic user information
	IRCOnly
	OsuUserStatsRequest // short, int32 (userid) - request to have a single BanchoHandleUserUpdate about an user. No idea what the first short is for.
	BanchoRestart
	OsuInvite
	BanchoInvite
	BanchoChannelListingComplete // Null - Finished sending channel names.
	OsuMatchChangePassword
	BanchoMatchChangePassword
	BanchoBanInfo // int32 - Number of seconds until the end of a silence.
	OsuSpecialMatchInfoRequest
	BanchoUserSilenced       // int32 - Broadcasted when a user is silenced, so that their messages are deleted.
	BanchoUserPresenceSingle // int32 - Broadcasted each time an user comes online.
	BanchoUserPresenceBundle // []int32 - max size: 512. value in the array is user ID.
	OsuUserPresenceRequest
	OsuUserPresenceRequestAll
	OsuUserToggleBlockNonFriendPM
	BanchoUserPMBlocked
	BanchoTargetIsSilenced
	BanchoVersionUpdateForced // force client update
	BanchoSwitchServer
	BanchoAccountRestricted
	BanchoRTX // pops up spooky message on your screen
	OsuMatchAbort
	BanchoSwitchTourneyServer
	OsuSpecialJoinMatchChannel  // force a client to join lobby (this is what OsuSQL uses afaik)
	OsuSpecialLeaveMatchChannel // force a client to leave lobby
)
