var RoboHydraHeadFilesystem = require("robohydra").heads.RoboHydraHeadFilesystem;
var RoboHydraHead = require("robohydra").heads.RoboHydraHead;
var RoboHydraHeadProxy = require("robohydra").heads.RoboHydraHeadProxy;

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
            })
        ]
    };
};
