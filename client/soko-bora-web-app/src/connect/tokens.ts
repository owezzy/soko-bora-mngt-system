import {createClientToken} from "./connect.module";
import {BasketService} from '../proto/basketspb/api_pb';
import {CustomersService} from "../proto/customerspb/api_pb";
import {OrderingService} from '../proto/orderingpb/api_pb';
import {PaymentsService} from '../proto/paymentspb/api_pb';
import {SearchService} from '../proto/searchpb/api_pb';

export const BasketGrpcService = createClientToken(BasketService);
export const CustomerGrpcService = createClientToken(CustomersService);
export const OrderingGrpcService = createClientToken(OrderingService);
export const PaymentsGrpcService = createClientToken(PaymentsService);
export const SearchGrpcService = createClientToken(SearchService);

// Additional client tokens representing Connect services could be added here
