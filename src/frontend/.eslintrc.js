/**
 * @file My eslint configuration, in JavaScript form.
 * @author Daniel Isaksen <d@duniel.no> (https://duniel.no)
 * @copyright Copyright (c) 2019, Daniel Isaksen. All rights reserved.
 */

module.exports = {
    env: {
        browser: true,
        node: true,
    },
    extends: "eslint:all",
    globals: {
        Atomics: "readonly",
        SharedArrayBuffer: "readonly",
    },
    parser: "babel-eslint",
    parserOptions: {
        ecmaVersion: 2022,
        sourceType: "module",
        ecmaFeatures: {
            impliedStrict: true,
            jsx: true,
        },
    },
    rules: {
        "max-lines": "off",
        "no-multi-spaces": "off",
        "object-curly-spacing": [ "warn", "always" ],
        "sort-imports": "off",
        "padded-blocks": "off",
        "max-statements": "off",
        "no-underscore-dangle": "off",
        "array-bracket-spacing": [ "warn", "always" ],
        "array-bracket-newline": [ "warn", "consistent" ],
        "camelcase": "off",
        "one-var": [ "warn", "consecutive" ],
        "space-before-function-paren": "off",
        "sort-vars": "off",
        "id-length": "off",
        "no-extra-parens": "off",
        "quote-props": "off",
        "no-sync": "off",
        "function-paren-newline": "off",
        "arrow-parens": [ "warn", "as-needed" ],
        "no-magic-numbers": "off",
        "object-property-newline": "off",
        "max-lines-per-function": "off",
        "key-spacing": "off",
        "array-element-newline": [ "warn", "consistent" ],
        "max-params": "off",
        "class-methods-use-this": "off",
        "max-depth": "off",
        "comma-dangle": [ "warn", "always-multiline" ],
        "no-unused-vars": "off",
        "capitalized-comments": "off",
        "linebreak-style": "off",
        "function-call-argument-newline": "off",
        "operator-linebreak": [ "warn", "before" ],
        "sort-keys": "off",
        "lines-around-comment": "off",
    },
};
