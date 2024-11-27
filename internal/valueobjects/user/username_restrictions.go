package user

import "github.com/kye-gregory/koicards-api/pkg/types/immutableslice"

var usernameBlacklist = immutableslice.NewImmutableSlice[string]([]string{
	"koicards","admin","root","moderator","system",
	"fck","fuck","fucker","bitch","bitches","asshole","arsehole","cunt","dick","slut",
	"bastard","shit","prick","cock","pussy","pussies","whore","hoe",
	"faggot","nigga","nigger","crap","twat","wanker","sex","porn","boobs","tits",
	"titties","dildo","vibrator","masturbate","masturbation","orgasm","cum","cumming",
	"ejaculate","anal","threesome","foursome","hooker","escort","nude","nudes",
	"stripper","lapdance","pegging","bdsm","dominatrix","fetish","incest","brothel",
	"pimp","hentai","blowjob","handjob","rimjob","doggy","missionary","racist",
	"homophobe","homophobic","bigot","murder","genocide","nazi","terrorist",
	"terrorism","islamophobe","jew","antisemite","retard","retarded",
})

var usernameWhitelist = immutableslice.NewImmutableSlice[string]([]string{
	"admiral","truck","batch","count","username",
})