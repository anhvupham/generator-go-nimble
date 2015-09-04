'use strict';
var yeoman = require('yeoman-generator');
var chalk = require('chalk');
var yosay = require('yosay');

module.exports = yeoman.generators.Base.extend({
    // constructor: function() {
    //     // generators.Base.apply(this, arguments);

    //     // This makes `appname` a required argument.
    //     // this.argument('appname', {
    //     //     type: String,
    //     //     required: true
    //     // });
    //     // And you can then access it later on this way; e.g. CamelCased
    //     // this.appname = _.camelCase(this.appname);
    // },

    prompting: function() {
        var done = this.async();

        // Have Yeoman greet the user.
        this.log(yosay(
            'Welcome to the perfect ' + chalk.red('GoNimble') + ' generator!'
        ));

        var prompts = [{
            type: 'input',
            name: 'port',
            message: 'Which port do you want to use ?',
            default: 8080
        }, {
            type: 'input',
            name: 'dbname',
            message: 'What is name of database ?',
            default: "Foo"
        }, {
            type: 'input',
            name: 'mongohost',
            message: 'What is MongoDB host address ?',
            default: "localhost"
        }, {
            type: 'input',
            name: 'redishost',
            message: 'What is Redis host address ?',
            default: "localhost"
        }, {
            type: 'input',
            name: 'redisport',
            message: 'What is Redis port ?',
            default: 6379
        }, {
            type: 'input',
            name: 'rediskey',
            message: 'What is Redis session key ?',
            default: "my-key"
        }];

        this.prompt(prompts, function(props) {
            this.props = props;
            // To access props later use this.props.someOption;

            done();
        }.bind(this));
    },

    writing: {
        app: function() {
            this.fs.copy(
                this.templatePath('auth.go'),
                this.destinationPath('src/auth.go')
            );
            this.fs.copy(
                this.templatePath('handlers.go'),
                this.destinationPath('src/handlers.go')
            );
            this.fs.copy(
                this.templatePath('main.go'),
                this.destinationPath('src/main.go')
            );
            this.fs.copy(
                this.templatePath('messages.go'),
                this.destinationPath('src/messages.go')
            );
            this.fs.copy(
                this.templatePath('models.go'),
                this.destinationPath('src/models.go')
            );
            this.fs.copy(
                this.templatePath('router.go'),
                this.destinationPath('src/router.go')
            );
            this.fs.copy(
                this.templatePath('routes.go'),
                this.destinationPath('src/routes.go')
            );
            this.fs.copy(
                this.templatePath('utils.go'),
                this.destinationPath('src/utils.go')
            );
            this.fs.copyTpl(
                this.templatePath('example_test.go'),
                this.destinationPath('src/example_test.go'), this.props
            );
            this.fs.copyTpl(
                this.templatePath('init.go'),
                this.destinationPath('src/init.go'), this.props
            );
            this.fs.copyTpl(
                this.templatePath('Dockerfile'),
                this.destinationPath('src/Dockerfile'), this.props
            );
        },

        projectfiles: function() {
            this.fs.copy(
                this.templatePath('editorconfig'),
                this.destinationPath('.editorconfig')
            );
        }
    },

    install: function() {
        this.spawnCommand('mkdir', ['bin']);
        this.spawnCommand('mkdir', ['pkg']);
    }
});
