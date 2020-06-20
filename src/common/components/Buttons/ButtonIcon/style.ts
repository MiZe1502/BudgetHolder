import { css } from 'emotion';

export { ButtonReset, ButtonDefault } from "../../../styles/buttons";

export const Button = css`
    width: 24px;
    height: 24px;
    background-color: transparent;

	&:hover,
	&:active,
    &:disabled {
		background-color: transparent;
	}
`;