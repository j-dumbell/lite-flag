/* tslint:disable */
/* eslint-disable */
/**
 * Liteflag API
 * API for managing feature flags.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import type { Configuration } from './configuration';
import type { AxiosPromise, AxiosInstance, RawAxiosRequestConfig } from 'axios';
import globalAxios from 'axios';
// Some imports not used depending on template conditions
// @ts-ignore
import { DUMMY_BASE_URL, assertParamExists, setApiKeyToObject, setBasicAuthToObject, setBearerAuthToObject, setOAuthToObject, setSearchParams, serializeDataIfNeeded, toPathString, createRequestFunction } from './common';
import type { RequestArgs } from './base';
// @ts-ignore
import { BASE_PATH, COLLECTION_FORMATS, BaseAPI, RequiredError, operationServerMap } from './base';

/**
 * 
 * @export
 * @interface ApiKey
 */
export interface ApiKey {
    /**
     * Unique identifier for the API key
     * @type {string}
     * @memberof ApiKey
     */
    'name': string;
    /**
     * The associated permissions of the key
     * @type {string}
     * @memberof ApiKey
     */
    'role': ApiKeyRoleEnum;
    /**
     * The API Key
     * @type {string}
     * @memberof ApiKey
     */
    'key': string;
}

export const ApiKeyRoleEnum = {
    Root: 'root',
    Admin: 'admin',
    Readonly: 'readonly'
} as const;

export type ApiKeyRoleEnum = typeof ApiKeyRoleEnum[keyof typeof ApiKeyRoleEnum];

/**
 * 
 * @export
 * @interface ApiKeyInput
 */
export interface ApiKeyInput {
    /**
     * Unique identifier for the API key
     * @type {string}
     * @memberof ApiKeyInput
     */
    'name': string;
    /**
     * The associated permissions of the key
     * @type {string}
     * @memberof ApiKeyInput
     */
    'role': ApiKeyInputRoleEnum;
}

export const ApiKeyInputRoleEnum = {
    Root: 'root',
    Admin: 'admin',
    Readonly: 'readonly'
} as const;

export type ApiKeyInputRoleEnum = typeof ApiKeyInputRoleEnum[keyof typeof ApiKeyInputRoleEnum];

/**
 * 
 * @export
 * @interface Flag
 */
export interface Flag {
    /**
     * Unique identifier for the feature flag
     * @type {string}
     * @memberof Flag
     */
    'key': string;
    /**
     * Type of the flag value
     * @type {string}
     * @memberof Flag
     */
    'type': FlagTypeEnum;
    /**
     * Whether or not the flag is public.
     * @type {boolean}
     * @memberof Flag
     */
    'isPublic': boolean;
    /**
     * 
     * @type {FlagValue}
     * @memberof Flag
     */
    'value': FlagValue;
}

export const FlagTypeEnum = {
    Boolean: 'boolean',
    String: 'string'
} as const;

export type FlagTypeEnum = typeof FlagTypeEnum[keyof typeof FlagTypeEnum];

/**
 * 
 * @export
 * @interface FlagInput
 */
export interface FlagInput {
    /**
     * Whether the feature flag is public or not
     * @type {boolean}
     * @memberof FlagInput
     */
    'isPublic': boolean;
    /**
     * Type of the flag value
     * @type {string}
     * @memberof FlagInput
     */
    'type': FlagInputTypeEnum;
    /**
     * 
     * @type {FlagValue}
     * @memberof FlagInput
     */
    'value': FlagValue;
}

export const FlagInputTypeEnum = {
    Boolean: 'boolean',
    String: 'string'
} as const;

export type FlagInputTypeEnum = typeof FlagInputTypeEnum[keyof typeof FlagInputTypeEnum];

/**
 * @type FlagValue
 * Value of the flag, must match the type
 * @export
 */
export type FlagValue = boolean | string;

/**
 * 
 * @export
 * @interface HealthResponse
 */
export interface HealthResponse {
    /**
     * Whether the database is healthy
     * @type {boolean}
     * @memberof HealthResponse
     */
    'database': boolean;
}

/**
 * APIKeysApi - axios parameter creator
 * @export
 */
export const APIKeysApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Delete an API key
         * @param {string} name Name of the API key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiKeysNameDelete: async (name: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'name' is not null or undefined
            assertParamExists('apiKeysNameDelete', 'name', name)
            const localVarPath = `/api-keys/{name}`
                .replace(`{${"name"}}`, encodeURIComponent(String(name)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Rotate an API key
         * @param {string} name Name of the API key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiKeysNameRotatePost: async (name: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'name' is not null or undefined
            assertParamExists('apiKeysNameRotatePost', 'name', name)
            const localVarPath = `/api-keys/{name}/rotate`
                .replace(`{${"name"}}`, encodeURIComponent(String(name)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Create a new API key
         * @param {ApiKeyInput} apiKeyInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiKeysPost: async (apiKeyInput: ApiKeyInput, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'apiKeyInput' is not null or undefined
            assertParamExists('apiKeysPost', 'apiKeyInput', apiKeyInput)
            const localVarPath = `/api-keys`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(apiKeyInput, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * APIKeysApi - functional programming interface
 * @export
 */
export const APIKeysApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = APIKeysApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Delete an API key
         * @param {string} name Name of the API key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiKeysNameDelete(name: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiKeysNameDelete(name, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['APIKeysApi.apiKeysNameDelete']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Rotate an API key
         * @param {string} name Name of the API key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiKeysNameRotatePost(name: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ApiKey>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiKeysNameRotatePost(name, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['APIKeysApi.apiKeysNameRotatePost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Create a new API key
         * @param {ApiKeyInput} apiKeyInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async apiKeysPost(apiKeyInput: ApiKeyInput, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<ApiKey>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.apiKeysPost(apiKeyInput, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['APIKeysApi.apiKeysPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * APIKeysApi - factory interface
 * @export
 */
export const APIKeysApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = APIKeysApiFp(configuration)
    return {
        /**
         * 
         * @summary Delete an API key
         * @param {string} name Name of the API key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiKeysNameDelete(name: string, options?: RawAxiosRequestConfig): AxiosPromise<void> {
            return localVarFp.apiKeysNameDelete(name, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Rotate an API key
         * @param {string} name Name of the API key
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiKeysNameRotatePost(name: string, options?: RawAxiosRequestConfig): AxiosPromise<ApiKey> {
            return localVarFp.apiKeysNameRotatePost(name, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Create a new API key
         * @param {ApiKeyInput} apiKeyInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        apiKeysPost(apiKeyInput: ApiKeyInput, options?: RawAxiosRequestConfig): AxiosPromise<ApiKey> {
            return localVarFp.apiKeysPost(apiKeyInput, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * APIKeysApi - object-oriented interface
 * @export
 * @class APIKeysApi
 * @extends {BaseAPI}
 */
export class APIKeysApi extends BaseAPI {
    /**
     * 
     * @summary Delete an API key
     * @param {string} name Name of the API key
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof APIKeysApi
     */
    public apiKeysNameDelete(name: string, options?: RawAxiosRequestConfig) {
        return APIKeysApiFp(this.configuration).apiKeysNameDelete(name, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Rotate an API key
     * @param {string} name Name of the API key
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof APIKeysApi
     */
    public apiKeysNameRotatePost(name: string, options?: RawAxiosRequestConfig) {
        return APIKeysApiFp(this.configuration).apiKeysNameRotatePost(name, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Create a new API key
     * @param {ApiKeyInput} apiKeyInput 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof APIKeysApi
     */
    public apiKeysPost(apiKeyInput: ApiKeyInput, options?: RawAxiosRequestConfig) {
        return APIKeysApiFp(this.configuration).apiKeysPost(apiKeyInput, options).then((request) => request(this.axios, this.basePath));
    }
}



/**
 * FlagsApi - axios parameter creator
 * @export
 */
export const FlagsApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary Retrieve all feature flags
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsGet: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/flags`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Delete a feature flag
         * @param {string} key Unique key of the feature flag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsKeyDelete: async (key: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'key' is not null or undefined
            assertParamExists('flagsKeyDelete', 'key', key)
            const localVarPath = `/flags/{key}`
                .replace(`{${"key"}}`, encodeURIComponent(String(key)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'DELETE', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Retrieve a single feature flag by key
         * @param {string} key Unique key of the feature flag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsKeyGet: async (key: string, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'key' is not null or undefined
            assertParamExists('flagsKeyGet', 'key', key)
            const localVarPath = `/flags/{key}`
                .replace(`{${"key"}}`, encodeURIComponent(String(key)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Update an existing feature flag
         * @param {string} key Unique key of the feature flag
         * @param {FlagInput} flagInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsKeyPut: async (key: string, flagInput: FlagInput, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'key' is not null or undefined
            assertParamExists('flagsKeyPut', 'key', key)
            // verify required parameter 'flagInput' is not null or undefined
            assertParamExists('flagsKeyPut', 'flagInput', flagInput)
            const localVarPath = `/flags/{key}`
                .replace(`{${"key"}}`, encodeURIComponent(String(key)));
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'PUT', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(flagInput, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
        /**
         * 
         * @summary Create a new feature flag
         * @param {Flag} flag 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsPost: async (flag: Flag, options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            // verify required parameter 'flag' is not null or undefined
            assertParamExists('flagsPost', 'flag', flag)
            const localVarPath = `/flags`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'POST', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            localVarHeaderParameter['Content-Type'] = 'application/json';

            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};
            localVarRequestOptions.data = serializeDataIfNeeded(flag, localVarRequestOptions, configuration)

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * FlagsApi - functional programming interface
 * @export
 */
export const FlagsApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = FlagsApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary Retrieve all feature flags
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async flagsGet(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Array<Flag>>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.flagsGet(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['FlagsApi.flagsGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Delete a feature flag
         * @param {string} key Unique key of the feature flag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async flagsKeyDelete(key: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<void>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.flagsKeyDelete(key, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['FlagsApi.flagsKeyDelete']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Retrieve a single feature flag by key
         * @param {string} key Unique key of the feature flag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async flagsKeyGet(key: string, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Flag>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.flagsKeyGet(key, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['FlagsApi.flagsKeyGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Update an existing feature flag
         * @param {string} key Unique key of the feature flag
         * @param {FlagInput} flagInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async flagsKeyPut(key: string, flagInput: FlagInput, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Flag>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.flagsKeyPut(key, flagInput, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['FlagsApi.flagsKeyPut']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
        /**
         * 
         * @summary Create a new feature flag
         * @param {Flag} flag 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async flagsPost(flag: Flag, options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<Flag>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.flagsPost(flag, options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['FlagsApi.flagsPost']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * FlagsApi - factory interface
 * @export
 */
export const FlagsApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = FlagsApiFp(configuration)
    return {
        /**
         * 
         * @summary Retrieve all feature flags
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsGet(options?: RawAxiosRequestConfig): AxiosPromise<Array<Flag>> {
            return localVarFp.flagsGet(options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Delete a feature flag
         * @param {string} key Unique key of the feature flag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsKeyDelete(key: string, options?: RawAxiosRequestConfig): AxiosPromise<void> {
            return localVarFp.flagsKeyDelete(key, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Retrieve a single feature flag by key
         * @param {string} key Unique key of the feature flag
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsKeyGet(key: string, options?: RawAxiosRequestConfig): AxiosPromise<Flag> {
            return localVarFp.flagsKeyGet(key, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Update an existing feature flag
         * @param {string} key Unique key of the feature flag
         * @param {FlagInput} flagInput 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsKeyPut(key: string, flagInput: FlagInput, options?: RawAxiosRequestConfig): AxiosPromise<Flag> {
            return localVarFp.flagsKeyPut(key, flagInput, options).then((request) => request(axios, basePath));
        },
        /**
         * 
         * @summary Create a new feature flag
         * @param {Flag} flag 
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        flagsPost(flag: Flag, options?: RawAxiosRequestConfig): AxiosPromise<Flag> {
            return localVarFp.flagsPost(flag, options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * FlagsApi - object-oriented interface
 * @export
 * @class FlagsApi
 * @extends {BaseAPI}
 */
export class FlagsApi extends BaseAPI {
    /**
     * 
     * @summary Retrieve all feature flags
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof FlagsApi
     */
    public flagsGet(options?: RawAxiosRequestConfig) {
        return FlagsApiFp(this.configuration).flagsGet(options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Delete a feature flag
     * @param {string} key Unique key of the feature flag
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof FlagsApi
     */
    public flagsKeyDelete(key: string, options?: RawAxiosRequestConfig) {
        return FlagsApiFp(this.configuration).flagsKeyDelete(key, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Retrieve a single feature flag by key
     * @param {string} key Unique key of the feature flag
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof FlagsApi
     */
    public flagsKeyGet(key: string, options?: RawAxiosRequestConfig) {
        return FlagsApiFp(this.configuration).flagsKeyGet(key, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Update an existing feature flag
     * @param {string} key Unique key of the feature flag
     * @param {FlagInput} flagInput 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof FlagsApi
     */
    public flagsKeyPut(key: string, flagInput: FlagInput, options?: RawAxiosRequestConfig) {
        return FlagsApiFp(this.configuration).flagsKeyPut(key, flagInput, options).then((request) => request(this.axios, this.basePath));
    }

    /**
     * 
     * @summary Create a new feature flag
     * @param {Flag} flag 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof FlagsApi
     */
    public flagsPost(flag: Flag, options?: RawAxiosRequestConfig) {
        return FlagsApiFp(this.configuration).flagsPost(flag, options).then((request) => request(this.axios, this.basePath));
    }
}



/**
 * HealthzApi - axios parameter creator
 * @export
 */
export const HealthzApiAxiosParamCreator = function (configuration?: Configuration) {
    return {
        /**
         * 
         * @summary API healthcheck
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        healthzGet: async (options: RawAxiosRequestConfig = {}): Promise<RequestArgs> => {
            const localVarPath = `/healthz`;
            // use dummy base URL string because the URL constructor only accepts absolute URLs.
            const localVarUrlObj = new URL(localVarPath, DUMMY_BASE_URL);
            let baseOptions;
            if (configuration) {
                baseOptions = configuration.baseOptions;
            }

            const localVarRequestOptions = { method: 'GET', ...baseOptions, ...options};
            const localVarHeaderParameter = {} as any;
            const localVarQueryParameter = {} as any;

            // authentication ApiKeyAuth required
            await setApiKeyToObject(localVarHeaderParameter, "X-API-KEY", configuration)


    
            setSearchParams(localVarUrlObj, localVarQueryParameter);
            let headersFromBaseOptions = baseOptions && baseOptions.headers ? baseOptions.headers : {};
            localVarRequestOptions.headers = {...localVarHeaderParameter, ...headersFromBaseOptions, ...options.headers};

            return {
                url: toPathString(localVarUrlObj),
                options: localVarRequestOptions,
            };
        },
    }
};

/**
 * HealthzApi - functional programming interface
 * @export
 */
export const HealthzApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = HealthzApiAxiosParamCreator(configuration)
    return {
        /**
         * 
         * @summary API healthcheck
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        async healthzGet(options?: RawAxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<HealthResponse>> {
            const localVarAxiosArgs = await localVarAxiosParamCreator.healthzGet(options);
            const localVarOperationServerIndex = configuration?.serverIndex ?? 0;
            const localVarOperationServerBasePath = operationServerMap['HealthzApi.healthzGet']?.[localVarOperationServerIndex]?.url;
            return (axios, basePath) => createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration)(axios, localVarOperationServerBasePath || basePath);
        },
    }
};

/**
 * HealthzApi - factory interface
 * @export
 */
export const HealthzApiFactory = function (configuration?: Configuration, basePath?: string, axios?: AxiosInstance) {
    const localVarFp = HealthzApiFp(configuration)
    return {
        /**
         * 
         * @summary API healthcheck
         * @param {*} [options] Override http request option.
         * @throws {RequiredError}
         */
        healthzGet(options?: RawAxiosRequestConfig): AxiosPromise<HealthResponse> {
            return localVarFp.healthzGet(options).then((request) => request(axios, basePath));
        },
    };
};

/**
 * HealthzApi - object-oriented interface
 * @export
 * @class HealthzApi
 * @extends {BaseAPI}
 */
export class HealthzApi extends BaseAPI {
    /**
     * 
     * @summary API healthcheck
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof HealthzApi
     */
    public healthzGet(options?: RawAxiosRequestConfig) {
        return HealthzApiFp(this.configuration).healthzGet(options).then((request) => request(this.axios, this.basePath));
    }
}



