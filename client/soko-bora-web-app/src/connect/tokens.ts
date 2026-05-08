import {createClientToken} from "./connect.module";
import {CustomersService} from "../proto/customerspb/api_pb";
import {SearchService} from '../proto/searchpb/api_pb';

export const CustomerGrpcService = createClientToken(CustomersService);
export const SearchGrpcService = createClientToken(SearchService);

// Additional client tokens representing Connect services could be added here
