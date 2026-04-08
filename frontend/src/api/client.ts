import createClient from 'openapi-fetch';
import type { paths } from './v1';

export const api = createClient<paths>({ baseUrl: '/api' });
