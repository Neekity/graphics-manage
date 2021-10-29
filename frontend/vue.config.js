const fs = require('fs');
const path = require('path');
module.exports = {
    outputDir: 'dist',
    transpileDependencies: [
        'vuetify'
    ],
    devServer: {
        hot: true,
        compress: true,
        disableHostCheck: true, //webpack4.0 开启热更新
    },
    configureWebpack: config => {

        function copyFileSync( source, target ) {

            var targetFile = target;

            // If target is a directory, a new file with the same name will be created
            if ( fs.existsSync( target ) ) {
                if ( fs.lstatSync( target ).isDirectory() ) {
                    targetFile = path.join( target, path.basename( source ) );
                }
            }

            fs.writeFileSync(targetFile, fs.readFileSync(source));
        }

        function copyFolderRecursiveSync( source, target ) {
            var files = [];

            // Check if folder needs to be created or integrated
            var targetFolder = path.join( target, path.basename( source ) );
            if ( !fs.existsSync( targetFolder ) ) {
                fs.mkdirSync( targetFolder );
            }

            // Copy
            if ( fs.lstatSync( source ).isDirectory() ) {
                files = fs.readdirSync( source );
                files.forEach( function ( file ) {
                    var curSource = path.join( source, file );
                    if ( fs.lstatSync( curSource ).isDirectory() ) {
                        copyFolderRecursiveSync( curSource, targetFolder );
                    } else {
                        copyFileSync( curSource, targetFolder );
                    }
                } );
            }
        }
        copyFolderRecursiveSync('/data/node_modules/tinymce/skins','/data/public/tinymce')
    }
}
