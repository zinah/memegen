var RoboHydraHeadFilesystem = require("robohydra").heads.RoboHydraHeadFilesystem;
var RoboHydraHead = require("robohydra").heads.RoboHydraHead;
var RoboHydraHeadProxy = require("robohydra").heads.RoboHydraHeadProxy;
var RoboHydraHeadStatic = require("robohydra").heads.RoboHydraHeadStatic;

exports.getBodyParts = function(conf) {
    return {    
        heads: [
		    new RoboHydraHeadFilesystem({
                mountPath: '/assets',
                documentRoot: 'assets'
            }),
            // a head to slow down the response before proxying to the default
            // test cat, detached by defult - attach to test timeouts
            new RoboHydraHead({
                detached: true,
                path: '/cat',
                handler: function(req, res, next) {
                    setTimeout(function() {
                        next(req, res);
                    }, 5000);
                }
            }),
            new RoboHydraHeadProxy({
                mountPath: '/cat',
                proxyTo: 'http://localhost:3000/assets/default_test_cat.png'
            }),
            new RoboHydraHeadStatic({
                path: '/joke',
                responses: [
                    {
                        contentType: 'application/json',
                        content: `
{
    "error": false,
    "category": "Programming",
    "type": "single",
    "joke": "I've got a really good UDP joke to tell you but I don’t know if you'll get it.",
    "flags": {
        "nsfw": false,
        "religious": false,
        "political": false,
        "racist": false,
        "sexist": false,
        "explicit": false
    },
    "id": 0,
    "safe": true,
    "lang": "en"
}
                    `
                },
                {
                    contentType: 'application/json',
                    content: `
{
    "error": false,
    "category": "Programming",
    "type": "twopart",
    "setup": "Why did the Python data scientist get arrested at customs?",
    "delivery": "She was caught trying to import pandas!",
    "flags": {
        "nsfw": false,
        "religious": false,
        "political": false,
        "racist": false,
        "sexist": false,
        "explicit": false
    },
    "id": 234,
    "safe": true,
    "lang": "en"
}
                `
            },
                {
                    contentType: 'application/json',
                    content: `
{
    "error": false,
    "category": "Programming",
    "type": "twopart",
    "setup": "why do python programmers wear glasses?",
    "delivery": "Because they can't C.",
    "flags": {
        "nsfw": false,
        "religious": false,
        "political": false,
        "racist": false,
        "sexist": false,
        "explicit": false
    },
    "safe": true,
    "id": 294,
    "lang": "en"
}
                    `
            }
                ]
            })
        ]
    };
};
