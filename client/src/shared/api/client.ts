import createFetchClient from 'openapi-fetch';
import createClient from 'openapi-react-query';
import type { paths } from '@/lib/api/openapi.d.ts';
import { API_HOST } from '@/shared/config/config';

export const fetchClient = createFetchClient<paths>({
  baseUrl: API_HOST,
});

export const $api = createClient(fetchClient);