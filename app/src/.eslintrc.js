module.exports = {
    "extends": [
        "eslint:recommended",
        "plugin:@typescript-eslint/eslint-recommended",
        "plugin:@typescript-eslint/recommended"
    ],
    "rules": {
        "indent": ["error", 2],
        "comma-dangle": ["error", "always-multiline"],
        "semi": ["error", "never"],
        "quotes": ["error", "single"],
    }
}